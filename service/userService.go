package service

import (
	"context"
	"encoding/json"
	"fmt"
	user "fold/protobuf/golang/grpc/user"
	"fold/service/model"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

type userProject struct {
	db       *gorm.DB
	ESClient *elasticsearch.Client
}

func (u userProject) LinkProject(ctx context.Context, request *user.LinkProjectRequest) (*user.LinkProjectResponse, error) {
	resp := new(user.LinkProjectResponse)
	for _, v := range request.ProjectIds {
		m := model.UserProject{
			UserId:    int(request.UserId),
			ProjectId: int(v),
		}
		u.db.Create(&m)
	}
	return resp, nil
}

func (u userProject) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	m := model.User{Name: request.Name}
	u.db.Create(&m)
	return m.ToProto(), nil
}

func getBody(field, Value string) string {
	body := ""
	switch field {
	case "username":
		body = fmt.Sprintf(
			`{"query": {"multi_match": {"query": "%s", "fields": ["users.name"]}}}`,
			Value)
	case "userid":
		body = fmt.Sprintf(
			`{"query": {"multi_match": {"query": "%s", "fields": ["users.id"]}}}`,
			Value)
	case "hashtag":
		body = fmt.Sprintf(
			`{"query": {"multi_match": {"query": "%s", "fields": ["hashtags.name"]}}}`,
			Value)
	case "slug":
		body = fmt.Sprintf(`{
  "query": {
    "fuzzy": {
      "slug": {
        "value": "%s",
        "fuzziness": 2,
        "prefix_length": 2,
        "max_expansions": 50
      }
    }
  }
}`, Value)
	case "desc":
		body = fmt.Sprintf(`{
  "query": {
    "fuzzy": {
      "description": {
        "value": "%s",
        "fuzziness": 2,
        "prefix_length": 3,
        "max_expansions": 50
      }
    }
  }
}`, Value)

	}
	return body
}

func (h userProject) GetUserProject(ctx context.Context, request *user.GetUserProjectRequest) (*user.GetUserProjectResponse, error) {
	resp := new(user.GetUserProjectResponse)
	body := getBody(request.GetField(), request.GetValue())
	res, err := h.ESClient.Search(
		h.ESClient.Search.WithContext(context.Background()),
		h.ESClient.Search.WithIndex("fold"),
		h.ESClient.Search.WithBody(strings.NewReader(body)),
		h.ESClient.Search.WithPretty(),
	)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Println(err)
		} else {
			log.Println(err)
		}
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	x := r["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, v := range x {
		d := v.(map[string]interface{})["_source"]
		val, err := json.Marshal(&d)
		if err != nil {
			return nil, err
		}
		var es elasticSource
		err = json.Unmarshal(val, &es)
		if err != nil {
			return nil, err
		}
		proj := user.Project{
			Id:          int32(es.Id),
			Name:        es.Name,
			Slug:        es.Slug,
			Description: es.Descriptions,
		}
		for _, ht := range es.Hashtags {
			proj.Hashtags = append(proj.Hashtags, &user.HashTag{
				Id:   int32(ht.Id),
				Name: ht.Name,
			})
		}
		for _, us := range es.Users {
			proj.User = append(proj.User, &user.UserObj{
				Id:   int32(us.Id),
				Name: us.Name,
			})
		}
		resp.Projects = append(resp.Projects, &proj)

	}

	return resp, nil
}

type elasticSource struct {
	Id           int              `json:"id"`
	Name         string           `json:"name"`
	Slug         string           `json:"slug"`
	Users        []elasticUser    `json:"users"`
	Hashtags     []elasticHashtag `json:"hashtags"`
	Descriptions string           `json:"description"`
}
type elasticUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type elasticHashtag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUserProject(db *gorm.DB, es *elasticsearch.Client) user.UserServer {
	return &userProject{db: db, ESClient: es}
}
