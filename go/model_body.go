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

type Body struct {

	UsergroupName string `json:"usergroup_name"`

	UserAuthorities string `json:"user_authorities"`

	Authenticate string `json:"authenticate,omitempty"`

	AccessIps []string `json:"access_ips"`

	CurrentUser string `json:"current_user"`
}
