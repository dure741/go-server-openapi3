/*
 * 用户组管理API
 *
 * 这是一个对用户组进行创建、修改、检索、和删除的API
 *
 * API version: 3.11
 * Contact: huanghui@zdns.cn
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"log"
	"net/http"

	//以下包用来转换yaml
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func UserGroupsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Handler\": \"UserGroupsGet arrived.\"}"))
}

func UserGroupsIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Handler\": \"UserGroupsIdDelete arrived.\"}"))
}

func UserGroupsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Handler\": \"UserGroupsPost arrived.\"}"))
}

func SwaggerJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	buffer, err := ioutil.ReadFile("./api/swagger.yaml")
	var jsonbody []byte
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		jsonbody, err = yaml.YAMLToJSON(buffer)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	w.Write(jsonbody)
}

// w.Write([]byte(`{
// 	"openapi": "3.0.1",
// 	"info": {
// 	  "title": "用户组管理API",
// 	  "description": "这是一个对用户组进行创建、修改、检索、和删除的API",
// 	  "contact": {
// 		"name": "huang hui",
// 		"url": "http://www.example.com/support",
// 		"email": "huanghui@zdns.cn"
// 	  },
// 	  "license": {
// 		"name": "Apache 2.0",
// 		"url": "http://www.apache.org/licenses/LICENSE-2.0.html"
// 	  },
// 	  "version": "3.11"
// 	},
// 	"paths": {
// 	  "/user_groups": {
// 		"get": {
// 		  "tags": [
// 			"userGroups"
// 		  ],
// 		  "summary": "获取用户组的信息，按页显示",
// 		  "parameters": [
// 			{
// 			  "name": "page_num",
// 			  "in": "query",
// 			  "description": "路径参数，请求的页数，int",
// 			  "required": false,
// 			  "style": "form",
// 			  "explode": true,
// 			  "schema": {
// 				"type": "integer"
// 			  }
// 			},
// 			{
// 			  "name": "page_size",
// 			  "in": "query",
// 			  "description": "路径参数，请求页数包含元素个数",
// 			  "required": false,
// 			  "style": "form",
// 			  "explode": true,
// 			  "schema": {
// 				"type": "integer"
// 			  }
// 			},
// 			{
// 			  "name": "name",
// 			  "in": "query",
// 			  "description": "查询参数，通过用户组名称查询用户组",
// 			  "required": false,
// 			  "style": "form",
// 			  "explode": true,
// 			  "schema": {
// 				"type": "string"
// 			  }
// 			}
// 		  ],
// 		  "responses": {
// 			"200": {
// 			  "description": "操作成功",
// 			  "content": {
// 				"application/json": {
// 				  "schema": {
// 					"$ref": "#/components/schemas/inline_response_200"
// 				  }
// 				}
// 			  }
// 			}
// 		  }
// 		},
// 		"post": {
// 		  "tags": [
// 			"userGroups"
// 		  ],
// 		  "summary": "新建一个用户组",
// 		  "description": "新建一个用户组，在body中设置参数",
// 		  "requestBody": {
// 			"content": {
// 			  "application/json": {
// 				"schema": {
// 				  "$ref": "#/components/schemas/body"
// 				}
// 			  }
// 			},
// 			"required": true
// 		  },
// 		  "responses": {
// 			"200": {
// 			  "description": "添加用户组成功",
// 			  "content": {
// 				"application/json": {
// 				  "schema": {
// 					"$ref": "#/components/schemas/inline_response_200_1"
// 				  }
// 				},
// 				"text/html": {
// 				  "schema": {
// 					"$ref": "#/components/schemas/inline_response_200_1"
// 				  }
// 				}
// 			  }
// 			},
// 			"401": {
// 			  "description": "Unauthorized",
// 			  "content": {
// 			  }
// 			},
// 			"422": {
// 			  "description": "Unprocessble Entity",
// 			  "content": {
// 				"application/json": {
// 				  "schema": {
// 					"$ref": "#/components/schemas/inline_response_422"
// 				  }
// 				},
// 				"text/html": {
// 				  "schema": {
// 					"$ref": "#/components/schemas/inline_response_422"
// 				  }
// 				}
// 			  }
// 			}
// 		  },
// 		  "x-codegen-request-body-name": "usergruop_name"
// 		}
// 	  },
// 	  "/user_groups/{id}": {
// 		"delete": {
// 		  "tags": [
// 			"userGroups"
// 		  ],
// 		  "summary": "删除特定用户组",
// 		  "description": "通过URL后缀参数{id}，删除特定用户组",
// 		  "parameters": [
// 			{
// 			  "name": "id",
// 			  "in": "path",
// 			  "required": true,
// 			  "style": "simple",
// 			  "explode": false,
// 			  "schema": {
// 				"type": "string"
// 			  }
// 			}
// 		  ],
// 		  "responses": {
// 			"200": {
// 			  "description": "OK",
// 			  "content": {
// 				"application/json": {
// 				  "schema": {
// 					"$ref": "#/components/schemas/inline_response_200_2"
// 				  }
// 				}
// 			  }
// 			}
// 		  }
// 		}
// 	  }
// 	},
// 	"components": {
// 	  "schemas": {
// 		"inline_response_200_1": {
// 		  "type": "object",
// 		  "properties": {
// 			"id": {
// 			  "type": "string"
// 			},
// 			"usergroup_name": {
// 			  "type": "string"
// 			},
// 			"user_authorities": {
// 			  "type": "string"
// 			},
// 			"authenticate": {
// 			  "type": "string"
// 			},
// 			"access_ips": {
// 			  "type": "array",
// 			  "items": {
// 				"type": "string"
// 			  }
// 			}
// 		  },
// 		  "example": {
// 			"access_ips": [
// 			  "access_ips",
// 			  "access_ips"
// 			],
// 			"usergroup_name": "usergroup_name",
// 			"authenticate": "authenticate",
// 			"id": "id",
// 			"user_authorities": "user_authorities"
// 		  }
// 		},
// 		"inline_response_200": {
// 		  "type": "object",
// 		  "properties": {
// 			"resources": {
// 			  "type": "array",
// 			  "items": {
// 				"$ref": "#/components/schemas/inline_response_200_resources"
// 			  }
// 			},
// 			"page_num": {
// 			  "type": "integer",
// 			  "format": "int64"
// 			},
// 			"page_size": {
// 			  "type": "integer",
// 			  "format": "int64"
// 			},
// 			"total_size": {
// 			  "type": "integer",
// 			  "format": "int64"
// 			}
// 		  },
// 		  "example": {
// 			"total_size": 1,
// 			"resources": [
// 			  {
// 				"access_ips": [
// 				  "access_ips",
// 				  "access_ips"
// 				],
// 				"usergroup_name": "usergroup_name",
// 				"authenticate": "authenticate",
// 				"id": "id",
// 				"user_authorities": "user_authorities"
// 			  },
// 			  {
// 				"access_ips": [
// 				  "access_ips",
// 				  "access_ips"
// 				],
// 				"usergroup_name": "usergroup_name",
// 				"authenticate": "authenticate",
// 				"id": "id",
// 				"user_authorities": "user_authorities"
// 			  }
// 			],
// 			"page_num": 0,
// 			"page_size": 6
// 		  }
// 		},
// 		"inline_response_200_2": {
// 		  "type": "object",
// 		  "properties": {
// 			"result": {
// 			  "type": "string"
// 			}
// 		  },
// 		  "example": {
// 			"result": "result"
// 		  }
// 		},
// 		"body": {
// 		  "required": [
// 			"access_ips",
// 			"current_user",
// 			"user_authorities",
// 			"usergroup_name"
// 		  ],
// 		  "type": "object",
// 		  "properties": {
// 			"usergroup_name": {
// 			  "type": "string"
// 			},
// 			"user_authorities": {
// 			  "type": "string"
// 			},
// 			"authenticate": {
// 			  "type": "string"
// 			},
// 			"access_ips": {
// 			  "type": "array",
// 			  "items": {
// 				"type": "string"
// 			  }
// 			},
// 			"current_user": {
// 			  "type": "string"
// 			}
// 		  }
// 		},
// 		"inline_response_422": {
// 		  "type": "object",
// 		  "properties": {
// 			"error": {
// 			  "type": "string"
// 			}
// 		  }
// 		},
// 		"inline_response_200_resources": {
// 		  "type": "object",
// 		  "properties": {
// 			"id": {
// 			  "type": "string",
// 			  "description": "jasoijeoif"
// 			},
// 			"usergroup_name": {
// 			  "type": "string"
// 			},
// 			"user_authorities": {
// 			  "type": "string"
// 			},
// 			"authenticate": {
// 			  "type": "string"
// 			},
// 			"access_ips": {
// 			  "type": "array",
// 			  "items": {
// 				"type": "string"
// 			  }
// 			}
// 		  },
// 		  "example": {
// 			"access_ips": [
// 			  "access_ips",
// 			  "access_ips"
// 			],
// 			"usergroup_name": "usergroup_name",
// 			"authenticate": "authenticate",
// 			"id": "id",
// 			"user_authorities": "user_authorities"
// 		  }
// 		}
// 	  }
// 	}
//   }`))
//}
