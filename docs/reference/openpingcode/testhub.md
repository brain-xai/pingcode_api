测试管理
测试库
用于查看和管理测试库的基本信息。
资源地址：{GET} https://rest_api_root/v1/testhub/libraries/{library_id}

资源属性
字段	类型	描述
id	String	
测试库的id。

url	String	
测试库的资源地址。

identifier	String	
测试库的标识。

name	String	
测试库的名称。

scope_type	String	
测试库的所属类型。

允许值: organization, user_group

scope_id	String	
测试库的所属id。

visibility	String	
测试库的可见性。

允许值: private, public

color	String	
测试库的主题色。

description	String	
测试库的描述。

members	Object[]	
测试库的成员列表。

created_at	Number	
测试库的创建时间。

created_by	Object	
测试库的创建人。

updated_at	Number	
测试库的更新时间。

updated_by	Object	
测试库的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#F693E7",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
引用数据示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "is_archived": 0,
    "is_deleted": 0
}
创建一个测试库
POST
https://rest_api_root/v1/testhub/libraries
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
scope_type可选	String	
测试库的所属类型。默认值organization。允许值分别表示企业可见和团队可见。

默认值: organization

允许值: organization, user_group

scope_id可选	String	
测试库的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。

name	String	
测试库的名称。

visibility可选	String	
测试库的可见范围。允许值分别表示组织全部成员可见和仅项目成员可见。

默认值: private

允许值: public, private

identifier	String	
测试库的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
测试库的描述。

members可选	Object[]	
测试库成员的列表。当测试库的scope_type变为user_group时，测试库成员必须是scope_id指定的团队内的成员；当scope_type为其他值时，测试库成员可以是任意成员或团队。

members.id	String	
企业成员或团队的id。

members.type	String	
测试库成员的类型。

允许值: user, user_group

请求示例：
{
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "测试库",
    "visibility": "private",
    "identifier": "CSK",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "type": "user"
        }
    ]
}
响应示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#F693E7",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
部分更新一个测试库
PATCH
https://rest_api_root/v1/testhub/libraries/{library_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

请求参数
字段	类型	描述
name可选	String	
测试库的名称。

identifier可选	String	
测试库的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
测试库的描述。

请求示例：
{
    "name": "测试库",
    "identifier": "CSK",
    "description": "这是一个测试库"
}
响应示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#F693E7",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
获取测试库列表
GET
https://rest_api_root/v1/testhub/libraries
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
name可选	String	
测试库的名称。

scope_type可选	String	
测试库的所属类型。

允许值: organization, user_group

scope_id可选	String	
测试库的所属id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47000",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
            "identifier": "CSK",
            "name": "测试库",
            "scope_type": "user_group",
            "scope_id": "63c8fb32729dee3334d96af7",
            "visibility": "private",
            "color": "#F693E7",
            "description": "这是一个测试库",
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
向测试库中添加一个成员
POST
https://rest_api_root/v1/testhub/libraries/{library_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

请求参数
字段	类型	描述
member可选	Object	
测试库的成员。

member.id	String	
企业成员或团队的id。

member.type	String	
成员的类型。

允许值: user, user_group

role_id可选	String	
角色的id。

请求示例：
{
    "member": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "type": "user"
    }
    "role_id": "6422711c3f12e6c1e46d40e6"
}
响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
部分更新一个测试库内的成员
PATCH
https://rest_api_root/v1/testhub/libraries/{library_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

member_id	String	
企业成员或团队的id。

请求参数
字段	类型	描述
role_id可选	String	
角色的id。管理员可以更改其他用户的角色，但非管理员用户无权更改其他用户的角色。

请求示例：
{
    "role_id": "6422711c3f12e6c1e46d40e6"
}
响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
获取测试库中的成员列表
GET
https://rest_api_root/v1/testhub/libraries/{library_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/63c8fb32729dee3334d96af7",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        }
    ]
}
在测试库中移除一个成员
DELETE
https://rest_api_root/v1/testhub/libraries/{library_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

member_id	String	
企业成员或团队的id。

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
向测试库中添加一个用例模块
POST
https://rest_api_root/v1/testhub/libraries/{library_id}/suites
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

请求参数
字段	类型	描述
name	String	
测试模块名称。测试模块为树形结构，同一层次下的名称不能重复。

parent_id可选	String	
父测试模块的id。

请求示例：
{
    "name": "登录",
    "parent_id": "5eb623f6a70571487ea46999"
}
响应示例：
{
    "id": "55714870a70ea4eb623f6700",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "登录",
    "parent": {
        "id": "5eb623f6a70571487ea46999",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
        "name": "用户"
    },
    "paths": "首页/用户"
}
部分更新一个测试库中用例模块
PATCH
https://rest_api_root/v1/testhub/libraries/{library_id}/suites/{suite_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

suite_id	String	
测试模块的id。

请求参数
字段	类型	描述
name可选	String	
测试模块名称。测试模块为树形结构，同一层次下的名称不能重复。

parent_id可选	String	
父测试模块的id。

请求示例：
{
    "name": "登录",
    "parent_id": "5eb623f6a70571487ea46999"
}
响应示例：
{
    "id": "55714870a70ea4eb623f6700",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "登录",
    "parent": {
        "id": "5eb623f6a70571487ea46999",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
        "name": "用户"
    },
    "paths": "首页/用户"
}
获取测试库中的用例模块列表
GET
https://rest_api_root/v1/testhub/libraries/{library_id}/suites
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

查询参数
字段	类型	描述
parent_id可选	String	
父测试模块的id。值可以是root或者某个模块id，当为空时，获取到所有的模块，当为root时，获取到所有的一级模块，当为某个模块id时，获取到该模块的直属子模块。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "55714870a70ea4eb623f6700",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "登录",
            "parent": {
                "id": "5eb623f6a70571487ea46999",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
                "name": "用户"
            },
            "paths": "首页/用户"
        }
    ]
}
在测试库中移除一个用例模块
请注意，删除一个模块会自动删除其所有的子模块。

DELETE
https://rest_api_root/v1/testhub/libraries/{library_id}/suites/{suite_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

suite_id	String	
测试模块的id。

响应示例：
{
    "id": "55714870a70ea4eb623f6701",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6701",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "注册",
    "parent": {
        "id": "5eb623f6a70571487ea46999",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
        "name": "用户"
    },
    "paths": "首页/用户"
}
用例
用于查看和管理测试用例的基本信息。
资源地址：{GET} https://rest_api_root/v1/testhub/cases/{case_id}

资源属性
字段	类型	描述
id	String	
用例的id。

url	String	
用例的资源地址。

library	Object	
用例所属的测试库。

identifier	String	
用例的标识。

title	String	
用例的标题。

level	String	
用例重要程度的名字。

short_id	String	
用例的短id。

html_url	String	
用例的html地址。

important_level	Object	
用例的重要程度。

suite	Object	
用例所属的测试模块。

state	Object	
用例的状态。

type	Object	
用例的类型。

maintenance	Object	
用例的维护人。

test_type	String	
用例的测试类型。允许值分别表示自动测试和手动测试。

允许值: automation, manual

description	String	
用例的描述。

precondition	String	
用例的前置条件。

properties	Object	
用例的自定义属性。

estimated_workload	Number	
用例的预估工时。

remaining_workload	Number	
用例的剩余工时。

steps	Object[]	
用例的步骤列表。

participants	Object[]	
用例的关注人列表。

public_image_token	String	
用例描述和自定义多行文本类型属性里获取图片资源token。获取一个用例和获取用例列表参数include_public_image_token值有效时返回。

created_at	Number	
用例的创建时间。

created_by	Object	
用例的创建人。

updated_at	Number	
用例的更新时间。

updated_by	Object	
用例的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": 8,
    "remaining_workload": 5,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa111",
            "description": "网页访问",
            "expected_value": null,
            "is_group": 1,
            "group_id": null
        },
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "在浏览器地址栏中输入https://pingcode.com",
            "expected_value": "成功访问PingCode官网",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa111"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "public_image_token": "IcF1VmJFF-UIi9lMU18m1NFFAruWXYx0ZAm90pJ2LGk",
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
引用数据示例：
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
创建一个用例
POST
https://rest_api_root/v1/testhub/cases
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
test_library_id	String	
测试库的id。

title	String	
测试用例的标题。

suite_id可选	String	
测试用例所属模块的id。

type_id可选	String	
测试用例类型的id。

important_level_id可选	String	
测试用例重要程度的id。

maintenance_id可选	String	
测试用例维护人的id。

participant_ids可选	String[]	
测试用例关注人的id列表。

properties可选	Object	
测试用例属性的键值对集合，需要注意的是，当前测试用例对应的属性方案需要包含这些测试用例属性，例如属性方案中包含prop_a和prop_b。

properties.prop_a可选	Object	
测试用例属性prop_a。

properties.prop_b可选	Object	
测试用例属性prop_b。

description可选	String	
测试用例的描述。

precondition可选	String	
测试用例的前置条件。

steps可选	Object[]	
测试用例的步骤列表。steps是整体更新的。

steps.step_id可选	String	
测试用例步骤的id。

steps.description可选	String	
测试用例步骤的描述。

steps.expected_value可选	String	
测试用例步骤的期望值。

steps.is_group可选	Number	
测试用例步骤类型是否为分组。

允许值: 0, 1

steps.group_id可选	String	
测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。

请求示例：
{
    "test_library_id": "5eb623f6a70571487ea47000",
    "title": "这是一个测试用例",
    "suite_id": "55714870a70ea4eb623f6700",
    "type_id": "5cf189b35de9c20620ad7153",
    "important_level_id": "57a109b35ae8c20630fd7256",
    "maintenance_id": "a0417f68e846aae315c85d24643678a9",
    "participant_ids": [
        "a0417f68e846aae315c85d24643678a9"
    ],
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "steps": [
        {
            "step_id": "5cdca524cade3a112b063071",
            "description": "UI测试",
            "is_group": 1
        },
        {
            "step_id": "5cdca524cade3a112b063072",
            "description": "在浏览器地址栏中输入https://pingcode.com",
            "expected_value": "成功访问PingCode官网",
            "is_group": 0,
            "group_id": "5cdca524cade3a112b063071"
        }
    ]
}
响应示例：
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": null,
    "remaining_workload": null,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "UI测试",
            "expected_value": null,
            "is_group": 1,
            "group_id": null
        },
        {
            "step_id": "524cad5edb06305cca2fa113",
            "description": "在浏览器地址栏中输入https://pingcode.com",
            "expected_value": "成功访问PingCode官网",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
批量创建用例
POST
https://rest_api_root/v1/testhub/cases/bulk
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
cases	Object[]	
创建单个测试用例必要参数的数组，数组内对象不能超过100个。

cases.test_library_id	String	
测试用例所属测试库id。

cases.title	String	
测试用例的标题，长度1～200有效字符。

cases.important_level_id可选	String	
测试用例重要程度的id。

cases.maintenance_id可选	String	
测试用例维护人的id。

cases.participant_ids可选	String[]	
测试用例关注人的id列表。

cases.properties可选	String	
测试用例属性的键值对集合。

cases.description可选	String	
测试用例的描述。

cases.precondition可选	String	
测试用例的前置条件。

cases.steps可选	Object[]	
测试用例的步骤列表。

cases.steps.step_id可选	String	
测试用例步骤的id。

cases.steps.description可选	String	
测试用例步骤的描述。

cases.steps.expected_value可选	String	
测试用例步骤的期望值。

cases.steps.is_group可选	Number	
测试用例步骤类型是否为分组。

允许值: 0, 1

cases.steps.group_id可选	String	
测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。

请求示例：
{
    "cases": [
        {
            "test_library_id": "5eb623f6a70571487ea47000",
            "title": "这是一个测试用例",
            "important_level_id": "57a109b35ae8c20630fd7256",
            "maintenance_id": "a0417f68e846aae315c85d24643678a9",
            "participant_ids": [
                "a0417f68e846aae315c85d24643678a9"
            ],
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "description": "测试用例的描述",
            "precondition": "前置条件的描述信息",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "UI测试",
                    "is_group": 1
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "在浏览器地址栏中输入https://pingcode.com",
                    "expected_value": "成功访问PingCode官网",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                }
            ]
        }
    ]
}
响应示例：
[
    {
        "state": "success",
        "case": {
            "id": "5edca524cad2fa112b06305d",
            "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305d",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "CSK-1",
            "title": "这是一个测试用例",
            "level": "P1",
            "short_id": "fdUw3C",
            "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
            "important_level": {
                "id": "57a109b35ae8c20630fd7256",
                "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
                "name": "P1"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "state": {
                "id": "686f62038668bbae4f4dd0c1",
                "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
                "name": "设计",
                "type": "pending"
            },
            "type": {
                "id": "5cf189b35de9c20620ad7153",
                "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
                "name": "功能测试"
            },
            "maintenance": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "test_type": "automation",
            "description": "测试用例的备注",
            "precondition": "前置条件的描述信息",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "estimated_workload": null,
            "remaining_workload": null,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "UI测试",
                    "expected_value": null,
                    "is_group": 1,
                    "group_id": null
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "在浏览器地址栏中输入https://pingcode.com",
                    "expected_value": "成功访问PingCode官网",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305d",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583293300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    }
]
部分更新一个用例
PATCH
https://rest_api_root/v1/testhub/cases/{case_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
case_id	String	
测试用例的id。

请求参数
字段	类型	描述
suite_id可选	String	
测试用例所属模块的id。

state_id可选	String	
测试用例状态的id。

type_id可选	String	
测试用例类型的id。

title可选	String	
测试用例的标题。

important_level_id可选	String	
测试用例重要程度的id。

maintenance_id可选	String	
测试用例维护人的id。

properties可选	Object	
测试用例属性的键值对集合。需要注意的是，当前测试用例对应的属性方案需要包含这些测试用例属性。

properties.prop_a可选	Object	
测试用例属性prop_a。

properties.prop_b可选	Object	
测试用例属性prop_b。

description可选	String	
测试用例的备注。

precondition可选	String	
测试用例的前置条件。

steps可选	Object[]	
测试用例的步骤列表。steps是整体更新的。

steps.step_id可选	String	
测试用例的步骤的id。如果step中不包含用于确定唯一性的step_id，那么这个step将被视为新建，并为之创建新的step_id。

steps.description可选	String	
测试用例的步骤的描述。

steps.expected_value可选	String	
测试用例的步骤的期望值。

steps.is_group可选	Number	
测试用例步骤类型是否为分组。

允许值: 0, 1

steps.group_id可选	String	
测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。

请求示例：
{
    "suite_id": "55714870a70ea4eb623f6700",
    "state_id": "686f62038668bbae4f4dd0c1",
    "type_id": "5cf189b35de9c20620ad7153",
    "title": "这是一个测试用例",
    "important_level_id": "57a109b35ae8c20630fd7256",
    "maintenance_id": "a0417f68e846aae315c85d24643678a9",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "UI测试",
            "is_group": 1
        },
        {
            "step_id": "524cad5edb06305cca2fa113",
            "description": "点击下一页按钮",
            "expected_value": "成功跳转至下一页",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        },
        {
            "description": "点击最后一页按钮",
            "expected_value": "成功跳转至最后一页",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        }
    ]
}
响应示例：
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": 8,
    "remaining_workload": 5,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "UI测试",
            "expected_value": null,
            "is_group": 1,
            "group_id": null
        },
        {
            "step_id": "524cad5edb06305cca2fa113",
            "description": "点击下一页按钮",
            "expected_value": "成功跳转至下一页",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        },
        {
            "step_id": "524cad5edb06305cca2fa114",
            "description": "点击最后一页按钮",
            "expected_value": "成功跳转至最后一页",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
批量部分更新用例
PATCH
https://rest_api_root/v1/testhub/cases/bulk
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
cases	Object[]	
部分更新测试用例的数组。

cases.case_id	String	
测试用例的id。

cases.state_id可选	String	
测试用例状态的id。

cases.type_id可选	String	
测试用例类型的id。

cases.title可选	String	
测试用例的标题。

cases.important_level_id可选	String	
测试用例重要程度的id。

cases.maintenance_id可选	String	
测试用例维护人的id。

cases.properties可选	Object[]	
测试用例属性的键值对集合，property中包含propertyKey、propertyValue和propertyType三个字段。需要注意的是，当前测试用例对应的属性方案需要包含这些测试用例属性。

cases.properties.prop_a可选	Object	
测试用例属性的自定义属性prop_a。

cases.properties.prop_b可选	Object	
测试用例属性的自定义属性prop_b。

cases.description可选	String	
测试用例的备注。

cases.precondition可选	String	
测试用例的前置条件。

cases.steps可选	Object[]	
测试用例的步骤列表。steps是整体更新的。

cases.steps.step_id可选	String	
测试用例的步骤的id。如果step中不包含用于确定唯一性的step_id，那么这个step将被视为新建，并为之创建新的step_id。

cases.steps.description可选	String	
测试用例的步骤的描述。

cases.steps.expected_value可选	String	
测试用例的步骤的期望值。

cases.steps.is_group可选	Number	
测试用例步骤类型是否为分组。

允许值: 0, 1

cases.steps.group_id可选	String	
测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。

请求示例：
{
    "cases": [
        {
            "case_id": "5edca524cad2fa112b06305c",
            "state_id": "686f62038668bbae4f4dd0c1",
            "type_id": "5cf189b35de9c20620ad7153",
            "title": "这是一个测试用例",
            "important_level_id": "57a109b35ae8c20630fd7256",
            "maintenance_id": "a0417f68e846aae315c85d24643678a9",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "description": "测试用例的描述",
            "precondition": "前置条件的描述信息",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "UI测试",
                    "is_group": 1
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "点击下一页按钮",
                    "expected_value": "成功跳转至下一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                },
                {
                    "description": "点击最后一页按钮",
                    "expected_value": "成功跳转至最后一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                }
            ]
        }
    ]
}
响应示例：
[
    {
        "state": "success",
        "case": {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "CSK-10",
            "title": "这是一个测试用例",
            "level": "P1",
            "short_id": "fdUw3C",
            "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
            "important_level": {
                "id": "57a109b35ae8c20630fd7256",
                "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
                "name": "P1"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "state": {
                "id": "686f62038668bbae4f4dd0c1",
                "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
                "name": "设计",
                "type": "pending"
            },
            "type": {
                "id": "5cf189b35de9c20620ad7153",
                "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
                "name": "功能测试"
            },
            "maintenance": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "test_type": "automation",
            "description": "测试用例的备注",
            "precondition": "前置条件的描述信息",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "estimated_workload": 8,
            "remaining_workload": 5,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "UI测试",
                    "expected_value": null,
                    "is_group": 1,
                    "group_id": null
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "点击下一页按钮",
                    "expected_value": "成功跳转至下一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                },
                {
                    "step_id": "524cad5edb06305cca2fa114",
                    "description": "点击最后一页按钮",
                    "expected_value": "成功跳转至最后一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 15832903300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    }
]
获取用例列表
GET
https://rest_api_root/v1/testhub/cases
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id可选	String	
测试库的id。

maintenance_id可选	String	
维护人的id。

state_ids可选	String	
状态的id，使用','分割，最多只能20个。

important_level_ids可选	String	
重要程度的id，使用','分割，最多只能20个。

tag_ids可选	String	
标签的id，使用','分割，最多只能20个。

created_between可选	String	
创建时间介于的时间范围，通过','分割起始时间。

updated_between可选	String	
更新时间介于的时间范围，通过','分割起始时间。

include_public_image_token可选	String	
包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.prop_b。

include_deleted可选	Boolean	
是否查询已删除的用例。该值默认为false。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "CSK-10",
            "title": "这是一个测试用例",
            "level": "P1",
            "short_id": "fdUw3C",
            "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
            "important_level": {
                "id": "57a109b35ae8c20630fd7256",
                "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
                "name": "P1"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "state": {
                "id": "686f62038668bbae4f4dd0c1",
                "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
                "name": "设计",
                "type": "pending"
            },
            "type": {
                "id": "5cf189b35de9c20620ad7153",
                "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
                "name": "功能测试"
            },
            "maintenance": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "test_type": "automation",
            "description": "测试用例的备注",
            "precondition": "前置条件的描述信息",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "estimated_workload": 8,
            "remaining_workload": 5,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa100",
                    "description": "UI测试",
                    "expected_value": null,
                    "is_group": 1,
                    "group_id": null
                },
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "在浏览器地址栏中输入https://pingcode.com",
                    "expected_value": "成功访问PingCode官网",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa100"
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "点击下一页按钮",
                    "expected_value": "成功跳转至下一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa100"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "IcF1VmJFF-UIi9lMU18m1NFFAruWXYx0ZAm90pJ2LGk",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583293300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
删除一个用例
DELETE
https://rest_api_root/v1/testhub/cases/{case_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
case_id	String	
测试用例的id。

响应示例：
{
    "id": "5edca524cad2fa112b06305d",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305d",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-1",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": 8,
    "remaining_workload": 5,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "在浏览器地址栏中输入https://pingcode.com",
            "expected_value": "成功访问PingCode官网",
            "is_group": 0,
            "group_id": null
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305d",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 1
}
获取用例模块列表
GET
https://rest_api_root/v1/testhub/case/suites?library_id={library_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "55714870a70ea4eb623f6700",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
            "name": "登录",
            "paths": "首页/用户"
        }
    ]
}
获取用例属性列表
GET
https://rest_api_root/v1/testhub/case/properties?library_id={library_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "environment",
            "url": "https://rest_api_root/v1/testhub/case_properties/environment",
            "name": "重现环境",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "测试"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "生产"
                }
            ]
        },
        {
            "id": "estimated_workload",
            "url": "https://rest_api_root/v1/testhub/case_properties/estimated_workload",
            "name": "预估工时",
            "type": "number",
            "options": null
        }
    ]
}
获取用例状态列表
GET
https://rest_api_root/v1/testhub/case/states?library_id={library_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 3,
    "values": [
        {
            "id": "686f62038668bbae4f4dd0c1",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
            "name": "设计",
            "type": "pending"
        },
        {
            "id": "686f62038668bbae4f4dd0c2",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c2",
            "name": "就绪",
            "type": "completed"
        },
        {
            "id": "686f62038668bbae4f4dd0c3",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c3",
            "name": "废弃",
            "type": "closed"
        }
    ]
}
获取用例类型列表
GET
https://rest_api_root/v1/testhub/case/types?library_id={library_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cf189b35de9c20620ad7153",
            "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
            "name": "功能测试"
        }
    ]
}
获取用例的执行历史列表
获取测试用例所有执行用例的最后一次执行结果。

GET
https://rest_api_root/v1/testhub/cases/{case_id}/histories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
case_id	String	
测试用例的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "65115f0939286e26e05a66db",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea/histories/65115f0939286e26e05a66db",
            "run": {
                "id": "547000eb6a70571487623fea",
                "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
                "status": "pass",
                "short_id": "Aq1u38yX",
                "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX"
            },
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870,
                "short_id": "7nNkViOD",
                "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "status": "pass",
            "executed_at": 1583290300,
            "executed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "pass",
                    "actual_value": null
                }
            ]
        }
    ]
}
计划
用于查看和管理测试计划的基本信息。
资源地址：{GET} https://rest_api_root/v1/testhub/libraries/{library_id}/plans/{plan_id}

资源属性
字段	类型	描述
id	String	
计划的id。

url	String	
计划的资源地址。

library	Object	
计划所属的测试库。

name	String	
计划的名称。

state	Object	
计划的状态。

start_at	Number	
计划的开始时间。

end_at	Number	
计划的结束时间。

short_id	String	
计划的短id。

html_url	String	
计划的html地址。

type	Object	
计划关联的类型。包括项目、发布和迭代。

project	Object	
计划关联的项目。

sprint	Object	
计划关联的迭代。

version	Object	
计划关联的发布。

assignee	Object	
计划的负责人。

summary	String	
计划测试报告的概要。

created_at	Number	
计划的创建时间。

created_by	Object	
计划的创建人。

updated_at	Number	
计划的更新时间。

updated_by	Object	
计划的更新人。

全量数据示例：
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "测试计划",
    "state": {
        "id": "652d0cb2b798f983d9c67c2b",
        "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2b",
        "name": "进行中",
        "type": "in_progress"
    },
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
    "type": {
        "id": "641d0ab2b998f883f9c67b2c",
        "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
        "name": "发布测试"
    },
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "sprint": null,
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623487ea47001f6a70571",
        "name": "1.0.0",
        "start_at": 1565255712,
        "end_at": 1565255879,
        "stage": {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "type": "pending",
            "color": "#FA8888"
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "summary": "测试报告的概要",
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
引用数据示例：
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "name": "测试计划",
    "status": "in_progress",
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
}
创建一个计划
POST
https://rest_api_root/v1/testhub/libraries/{library_id}/plans
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

请求参数
字段	类型	描述
name	String	
测试计划的名称。名称在一个测试库里唯一。

type_id	String	
测试计划类型的id。

start_at	Number	
测试计划的开始时间。

end_at	Number	
测试计划的结束时间。

assignee_id	String	
测试计划负责人的id。

project_id可选	String	
项目的id。该字段在 sprint_id 或 version_id 有值时必填。

sprint_id可选	String	
迭代的id。该字段仅在 type_id 代表迭代测试时有效。

version_id可选	String	
发布的id。该字段仅在 type_id 代表发布测试时有效。

请求示例：
{
    "name": "测试计划",
    "type_id": "641d0ab2b998f883f9c67b2f",
    "project_id": "5eb623f6a70571487ea41919",
    "version_id": "641d0ab2b998f883f9c67b2c",
    "start_at": 1589791860,
    "end_at": 1589791870,
    "assignee_id": "a0417f68e846aae315c85d24643678a9"
}
响应示例：
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "测试计划",
    "state": {
        "id": "652d0cb2b798f983d9c67c2b",
        "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2c",
        "name": "进行中",
        "type": "in_progress"
    },
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
    "type": {
        "id": "641d0ab2b998f883f9c67b2c",
        "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
        "name": "发布测试"
    },
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "sprint": null,
    "version": {
        "id": "5eb623f6a70571487ea47001",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623f6a70571487ea47001",
        "name": "1.0.0",
        "start_at": 1565255712,
        "end_at": 1565255879,
        "stage": {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "type": "pending",
            "color": "#FA8888"
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "summary": "",
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
部分更新一个计划
PATCH
https://rest_api_root/v1/testhub/libraries/{library_id}/plans/{plan_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

plan_id	String	
测试计划的id。

请求参数
字段	类型	描述
name可选	String	
测试计划的名称。名称在一个测试库里唯一。

type_id可选	String	
测试计划类型的id。指定测试计划类型时，建议同时指定对应的 sprint_id 或 version_id。

project_id可选	String	
项目的id。

sprint_id可选	String	
迭代的id。该字段仅在测试计划类型为迭代测试时有效。

version_id可选	String	
发布的id。该字段仅在测试计划类型为发布测试时有效。

start_at可选	Number	
测试计划的开始时间。

end_at可选	Number	
测试计划的结束时间。

assignee_id可选	String	
测试计划负责人的id。

state_id可选	String	
测试计划状态的id。

summary可选	String	
测试报告的概要信息。

请求示例：
{
    "name": "测试计划",
    "type_id": "641d0ab2b998f883f9c67b2c",
    "project_id": "5eb623f6a70571487ea41919",
    "version_id": "5eb623487ea47001f6a70571",
    "start_at": 1589791860,
    "end_at": 1589791870,
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "state_id": "652d0cb2b798f983d9c67c2b",
    "summary": "测试报告的概要"
}
响应示例：
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "测试计划",
    "state": {
        "id": "652d0cb2b798f983d9c67c2b",
        "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2c",
        "name": "进行中",
        "type": "in_progress"
    },
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
    "type": {
        "id": "641d0ab2b998f883f9c67b2c",
        "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
        "name": "发布测试"
    },
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "sprint": null,
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623487ea47001f6a70571",
        "name": "1.0.0",
        "start_at": 1565255712,
        "end_at": 1565255879,
        "stage": {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "type": "pending",
            "color": "#FA8888"
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "summary": "测试报告的概要",
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
获取计划列表
GET
https://rest_api_root/v1/testhub/libraries/{library_id}/plans
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

查询参数
字段	类型	描述
name可选	String	
测试计划名称。

created_between可选	String	
创建时间介于的时间范围，通过','分割起始时间。

updated_between可选	String	
更新时间介于的时间范围，通过','分割起始时间。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb6a70571487623fea47000",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "测试计划",
            "state": {
                "id": "652d0cb2b798f983d9c67c2b",
                "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2b",
                "name": "进行中",
                "type": "in_progress"
            },
            "start_at": 1589791860,
            "end_at": 1589791870,
            "short_id": "7nNkViOD",
            "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
            "type": {
                "id": "641d0ab2b998f883f9c67b2c",
                "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
                "name": "发布测试"
            },
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "sprint": null,
            "version": {
                "id": "5eb623487ea47001f6a70571",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623487ea47001f6a70571",
                "name": "1.0.0",
                "start_at": 1565255712,
                "end_at": 1565255879,
                "stage": {
                    "id": "5f44a8f8bb347b14b49624a1",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                    "name": "未开始",
                    "type": "pending",
                    "color": "#FA8888"
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "summary": "测试报告的概要",
            "created_at": 1565366200,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1565366200,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
获取计划类型列表
GET
https://rest_api_root/v1/testhub/libraries/{library_id}/plan_types
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "642f765b6950bc66cfa82f05",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/642f765b6950bc66cfa82f05",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "普通测试"
        }
    ]
}
创建一个执行用例
POST
https://rest_api_root/v1/testhub/runs
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
library_id	String	
测试库的id。

plan_id	String	
测试计划的id。

case_id	String	
测试用例的id。

executor_id可选	String	
执行人的id。

请求示例：
{
    "library_id": "5eb623f6a70571487ea47000",
    "plan_id": "5eb6a70571487623fea47000",
    "case_id": "5edca524cad2fa112b06305c",
    "executor_id": "a0417f68e846aae315c85d24643678a9"
}
响应示例：
{
    "id": "547000eb6a70571487623fea",
    "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
    "status": "not_start",
    "short_id": "Aq1u38yX",
    "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "plan": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870,
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
    },
    "case": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "short_id": "fdUw3C",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "test_type": "automation",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "latest_executed_status": {
        "id": "68d117800d5dd2484a198265",
        "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
        "name": "未测"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "remark": null,
    "executor": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status": "not_start",
            "actual_value": null
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
批量创建执行用例
POST
https://rest_api_root/v1/testhub/runs/bulk
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
runs	Object[]	
创建单个执行用例必要参数的数组。数组长度不超过100。

runs.library_id	String	
执行用例所属测试库的id。

runs.plan_id	String	
执行用例所属测试计划的id。

runs.case_id	String	
测试用例的id。

runs.executor_id可选	String	
执行人的id。

请求示例：
{
    "runs": [
        {
            "library_id": "5edca524cad2fa112b06305a",
            "plan_id": "5edca524cad2fa112b06305b",
            "case_id": "5edca524cad2fa112b06305c",
            "executor_id": "a0417f68e846aae315c85d24643678a9"
        },
        {
            "library_id": "5edca524cad2fa112b06306a",
            "plan_id": "5edca524cad2fa112b06306b",
            "case_id": "5edca524cad2fa112b06306c",
            "executor_id": "a0417f68e846aae315c85d24643678b9"
        }
    ]
}
响应示例：
[
    {
        "state": "success",
        "run": {
            "id": "547000eb6a70571487623fea",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "remark": "执行用例执行结果的备注",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": null
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    },
    {
        "state": "failure",
        "run": {
            "library_id": "5edca524cad2fa112b06305a",
            "plan_id": "5edca524cad2fa112b06305b",
            "case_id": "5edca524cad2fa112b06305d",
            "executor_id": "a0417f68e846aae315c85d24643678a9"
        },
        "message": "创建失败或已创建"
    }
]
全量更新一个执行用例
PUT
https://rest_api_root/v1/testhub/runs/{run_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
run_id	String	
执行用例的id。

请求参数
字段	类型	描述
status_id	String	
执行用例执行结果的id。

remark可选	String	
执行用例执行结果的备注。

steps	Object[]	
执行用例步骤的列表。

steps.step_id	String	
执行用例步骤的id。

steps.status_id	String	
执行用例步骤执行结果的id。

steps.actual_value可选	String	
执行用例步骤的实际值。

executor_id可选	String	
执行人的id。不传默认执行人为空。

请求示例：
{
    "status_id": "68d117800d5dd2484a198261",
    "remark": "执行用例执行结果的备注",
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status_id": "68d117800d5dd2484a198261",
            "actual_value": "正常访问PingCode官网"
        }
    ],
    "executor_id": "a0417f68e846aae315c85d24643678a9"
}
响应示例：
{
    "id": "547000eb6a70571487623fea",
    "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
    "status": "pass",
    "short_id": "Aq1u38yX",
    "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "plan": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870,
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
    },
    "case": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "short_id": "fdUw3C",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "test_type": "automation",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "latest_executed_status": {
        "id": "68d117800d5dd2484a198261",
        "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
        "name": "通过"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "remark": "执行用例执行结果的备注",
    "executor": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status": "pass",
            "actual_value": null
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
部分更新一个执行用例
PATCH
https://rest_api_root/v1/testhub/runs/{run_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
run_id	String	
执行用例的id。

请求参数
字段	类型	描述
status_id	String	
执行用例执行结果的id。

remark可选	String	
执行用例执行结果的备注。

steps可选	Object[]	
执行用例步骤的列表。

steps.step_id	String	
执行用例步骤的id。

steps.status_id	String	
执行用例步骤执行结果的id。

steps.actual_value可选	String	
执行用例步骤的实际值。

executor_id可选	String	
执行人的id。不传默认执行人为执行用例的创建人。

请求示例：
{
    "status_id": "68d117800d5dd2484a198265",
    "remark": "执行用例执行结果的备注",
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status_id": "68d117800d5dd2484a198265",
            "actual_value": "正常访问PingCode官网"
        }
    ],
    "executor_id": "a0417f68e846aae315c85d24643678a9"
}
响应示例：
{
    "id": "547000eb6a70571487623fea",
    "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
    "status": "not_start",
    "short_id": "Aq1u38yX",
    "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "plan": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870,
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
    },
    "case": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "test_type": "automation",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "latest_executed_status": {
        "id": "68d117800d5dd2484a198265",
        "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
        "name": "未测"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "remark": "执行用例执行结果的备注",
    "executor": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status": "not_start",
            "actual_value": null
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
批量部分更新执行用例
PATCH
https://rest_api_root/v1/testhub/runs/bulk
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
runs	Object[]	
部分更新单个执行用例必要参数的数组。

runs.run_id	String	
执行用例的id。

runs.status_id	String	
执行用例执行结果的id。

runs.remark可选	String	
执行用例执行结果的备注。

runs.steps可选	Object[]	
执行用例的步骤列表。

runs.steps.step_id	String	
执行用例步骤的id。

runs.steps.status_id	String	
执行用例步骤执行结果的id。

runs.steps.actual_value可选	String	
执行用例步骤的实际值。

runs.executor_id可选	String	
执行人的id。

请求示例：
{
    "runs": [
        {
            "run_id": "5edca524cad2fa112b06305c",
            "status_id": "68d117800d5dd2484a198265",
            "remark": "执行用例执行结果的备注",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status_id": "68d117800d5dd2484a198265",
                    "actual_value": "正常访问PingCode官网"
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "status_id": "68d117800d5dd2484a198265",
                    "actual_value": "不正常访问PingCode官网"
                }
            ],
            "executor_id": "a0417f68e846aae315c85d24643678a9"
        },
        {
            "run_id": "5edca524cad2fa112b06305d",
            "status_id": "68d117800d5dd2484a198265",
            "remark": "执行用例执行结果的备注",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa114",
                    "status_id": "68d117800d5dd2484a198265",
                    "actual_value": "正常访问PingCode官网"
                },
                {
                    "step_id": "524cad5edb06305cca2fa114",
                    "status_id": "68d117800d5dd2484a198265",
                    "actual_value": "不正常访问PingCode官网"
                }
            ],
            "executor_id": "a0417f68e846aae315c85d24643678a8"
        }
    ]
}
响应示例：
[
    {
        "state": "success",
        "run": {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "remark": null,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": "正常访问PingCode官网"
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583299300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    },
    {
        "state": "success",
        "run": {
            "id": "5edca524cad2fa112b06305d",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "remark": "执行用例执行结果的备注",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": "正常访问PingCode官网"
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583299300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    }
]
批量操作执行用例
POST
https://rest_api_root/v1/testhub/libraries/{library_id}/plans/{plan_id}/runs/bulk
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
library_id	String	
测试库的id。

plan_id	String	
测试计划的id。

请求参数
字段	类型	描述
inserts可选	Object[]	
需要批量创建的执行用例。该参数是一个对象数组（数组内对象不得超过50个）。

inserts.case_id	String	
测试用例id。

inserts.executor_id可选	String	
执行人id。

updates可选	Object[]	
需要批量更新的执行用例。该参数是一个对象数组（数组内对象不得超过50个）。

updates.run_id	String	
执行用例的id。

updates.status_id	String	
执行用例执行结果的id。

updates.steps可选	Object[]	
执行用例步骤的数组。

updates.steps.step_id	String	
执行用例步骤的id。

updates.steps.status_id	String	
执行用例步骤执行结果的id。

updates.steps.actual_value可选	String	
执行用例步骤的实际值。

updates.executor_id可选	String	
执行人的id。

deletes可选	String[]	
需要批量删除的执行用例。该参数是一个执行用例id的数组（数组内id不得超过50个）。

请求示例：
{
    "inserts": [
        {
            "case_id": "5edca524cad2fa112b06305c",
            "executor_id": "a0417f68e846aae315c85d24643678a9"
        }
    ],
    "updates": [
        {
            "run_id": "547000eb6a70571487623fea",
            "status_id": "68d117800d5dd2484a198265",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status_id": "68d117800d5dd2484a198265",
                    "actual_value": "正常访问PingCode官网"
                }
            ],
            "executor_id": "a0417f68e846aae315c85d24643678a9"
        }
    ],
    "deletes": [
        "547000eb6a70571487623fea"
    ]
}
响应示例：
{
    "inserts": 1,
    "updates": 1,
    "deletes": 1
}
获取执行用例列表
GET
https://rest_api_root/v1/testhub/runs
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
plan_id可选	String	
测试计划的id。

case_id可选	String	
测试用例的id。

suite_id可选	String	
测试模块的id。

status_id可选	String	
执行用例执行结果的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "547000eb6a70571487623fea",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870,
                "short_id": "7nNkViOD",
                "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "remark": "执行用例执行结果的备注",
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": null
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
获取执行用例执行结果列表
GET
https://rest_api_root/v1/testhub/run/statuses?library_id={library_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id	String	
测试库的id。

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 5,
    "values": [
        {
            "id": "68d117800d5dd2484a198261",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
            "name": "通过"
        },
        {
            "id": "68d117800d5dd2484a198262",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198262",
            "name": "受阻"
        },
        {
            "id": "68d117800d5dd2484a198263",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198263",
            "name": "失败"
        },
        {
            "id": "68d117800d5dd2484a198264",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198264",
            "name": "跳过"
        },
        {
            "id": "68d117800d5dd2484a198265",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
            "name": "未测"
        }
    ]
}
获取执行用例的结果记录
GET
https://rest_api_root/v1/testhub/runs/{run_id}/histories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
run_id	String	
执行用例的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "65115f0939286e26e05a66db",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea/histories/65115f0939286e26e05a66db",
            "run": {
                "id": "547000eb6a70571487623fea",
                "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
                "status": "pass",
                "short_id": "Aq1u38yX",
                "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX"
            },
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870,
                "short_id": "7nNkViOD",
                "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "executed_status": {
                "id": "68d117800d5dd2484a198261",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
                "name": "通过"
            },
            "remark": "执行用例执行结果的备注",
            "executed_at": 1583290300,
            "executed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "pass",
                    "actual_value": null
                }
            ]
        }
    ]
}
测试配置中心
用于查看和管理测试用例属性。

用例配置
用于查看和管理测试用例属性。

获取全部用例状态列表
GET
https://rest_api_root/v1/testhub/case_states
权限: 企业令牌/用户令牌

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 3,
    "values": [
        {
            "id": "686f62038668bbae4f4dd0c1",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
            "name": "设计",
            "type": "pending"
        },
        {
            "id": "686f62038668bbae4f4dd0c2",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c2",
            "name": "就绪",
            "type": "completed"
        },
        {
            "id": "686f62038668bbae4f4dd0c3",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c3",
            "name": "废弃",
            "type": "closed"
        }
    ]
}
获取全部用例类型列表
GET
https://rest_api_root/v1/testhub/case_types
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cf189b35de9c20620ad7153",
            "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
            "name": "功能测试"
        }
    ]
}
获取全部重要程度列表
GET
https://rest_api_root/v1/testhub/case_important_levels
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "57a109b35ae8c20630fd7256",
            "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
            "name": "P11",
            "color": "#56ABFB"
        }
    ]
}
创建一个用例属性
POST
https://rest_api_root/v1/testhub/case_properties
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
用例属性的名称。在一个企业中这个名称是唯一的。

type	String	
用例属性的类型。

允许值: text, textarea, select, multi_select, cascade_select, cascade_multi_select, member, members, date, number, progress, rate, link

options可选	Object[]	
选择项列表。当用例属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例：
{
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "text": "严重"
        },
        {
            "text": "一般"
        }
    ]
}
响应示例：
{
    "id": "severity",
    "url": "https://rest_api_root/v1/testhub/case_properties/severity",
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true
}
部分更新一个用例属性
PATCH
https://rest_api_root/v1/testhub/case_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_id	String	
用例属性的id。

请求参数
字段	类型	描述
name可选	String	
用例属性的名称。在一个企业中这个名称是唯一的。

options可选	Object[]	
选择项列表。options是整体更新的。当用例属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例：
{
    "name": "严重程度-update",
    "options": [
        {
            "id": "5efb1859110533727a82c603",
            "text": "严重-update"
        },
        {
            "text": "一般"
        }
    ]
}
响应示例：
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/testhub/case_properties/severity",
    "name": "严重程度-update",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重-update"
        },
        {
            "_id": "5efb1859110533727a82c624",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true
}
获取全部用例属性列表
GET
https://rest_api_root/v1/testhub/case_properties
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "environment",
            "url": "https://rest_api_root/v1/testhub/case_properties/environment",
            "name": "重现环境",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "测试"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "生产"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "estimated_workload",
            "url": "https://rest_api_root/v1/testhub/case_properties/estimated_workload",
            "name": "预估工时",
            "type": "number",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
获取用例属性方案列表
GET
https://rest_api_root/v1/testhub/case_property_plans
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
library_id可选	String	
测试库的id。查询开启本地配置的方案时传入。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5f8a21f18ef715265de90c21",
            "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
            "category": "library",
            "host": "case",
            "library": null
        },
        {
            "id": "5f8a21f18ef715265de90c22",
            "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c22",
            "category": "library",
            "host": "case",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向属性方案中添加一个用例属性
POST
https://rest_api_root/v1/testhub/case_property_plans/{property_plan_id}/case_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
测试用例属性方案的id。

请求参数
字段	类型	描述
property_id	String	
测试用例属性的id。

请求示例：
{
    "property_id": "environment"
}
响应示例：
{
    "id": "environment",
    "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21/case_properties/environment",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
        "category": "library",
        "host": "case"
    },
    "property": {
        "id": "environment",
        "url": "https://rest_api_root/v1/testhub/case_properties/environment",
        "name": "重现环境",
        "type": "select",
        "options": [
            {
                "_id": "5efb1859110533727a82c603",
                "text": "test"
            },
            {
                "_id": "5efb1859110533727a82c604",
                "text": "beta"
            },
            {
                "_id": "5efb1859110533727a82c605",
                "text": "rc"
            }
        ]
    }
}
获取属性方案中的用例属性列表
GET
https://rest_api_root/v1/testhub/case_property_plans/{property_plan_id}/case_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
测试用例属性方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "environment",
            "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21/case_properties/environment",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
                "category": "library",
                "host": "case"
            },
            "property": {
                "id": "environment",
                "url": "https://rest_api_root/v1/testhub/case_properties/environment",
                "name": "重现环境",
                "type": "select",
                "options": [
                    {
                        "_id": "5efb1859110533727a82c603",
                        "text": "test"
                    },
                    {
                        "_id": "5efb1859110533727a82c604",
                        "text": "beta"
                    },
                    {
                        "_id": "5efb1859110533727a82c605",
                        "text": "rc"
                    }
                ]
            }
        },
        {
            "id": "estimated_workload",
            "url": "https://rest_api_root/v1/testhub/property_plans/5f8a21f18ef715265de90c21/properties/estimated_workload",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/testhub/property_plans/5f8a21f18ef715265de90c21",
                "category": "library",
                "host": "case"
            },
            "property": {
                "id": "estimated_workload",
                "url": "https://rest_api_root/v1/testhub/case_properties/estimated_workload",
                "name": "预估工时",
                "type": "number",
                "options": null
            }
        }
    ]
}
在属性方案中移除一个用例属性
DELETE
https://rest_api_root/v1/testhub/case_property_plans/{property_plan_id}/case_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
测试用例属性方案的id。

property_id	String	
测试用例属性的id。

响应示例：
{
    "id": "environment",
    "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21/case_properties/environment",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
        "category": "library",
        "host": "case"
    },
    "property": {
        "id": "environment",
        "url": "https://rest_api_root/v1/testhub/case_properties/environment",
        "name": "重现环境",
        "type": "select",
        "options": [
            {
                "_id": "5efb1859110533727a82c603",
                "text": "test"
            },
            {
                "_id": "5efb1859110533727a82c604",
                "text": "beta"
            },
            {
                "_id": "5efb1859110533727a82c605",
                "text": "rc"
            }
        ]
    }
}
执行用例配置
用于查看和管理执行用例属性。

获取全部执行用例执行结果列表
GET
https://rest_api_root/v1/testhub/run_statuses
权限: 企业令牌/用户令牌

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 5,
    "values": [
        {
            "id": "68d117800d5dd2484a198261",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
            "name": "通过",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198262",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198262",
            "name": "受阻",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198263",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198263",
            "name": "失败",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198264",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198264",
            "name": "跳过",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198265",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
            "name": "未测",
            "is_system": "true"
        }
    ]
}
计划配置
用于查看和管理计划属性。

获取全部计划状态列表
GET
https://rest_api_root/v1/testhub/plan_states
权限: 企业令牌/用户令牌

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 4,
    "values": [
        {
            "id": "686f62038668bbae4f4dd0ca",
            "url": "http://rest_api_root/v1/testhub/plan_states/686f62038668bbae4f4dd0ca",
            "name": "未开始",
            "type": "pending",
            "is_system": true
        },
        {
            "id": "652d0cb2b798f983d9c67c2b",
            "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2b",
            "name": "进行中",
            "type": "in_progress",
            "is_system": true
        },
        {
            "id": "652d0cb2b798f983d9c67c2d",
            "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2d",
            "name": "已完成",
            "type": "completed",
            "is_system": true
        },
        {
            "id": "652d0cb2b798f983d9c67c2e",
            "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2e",
            "name": "已终止",
            "type": "completed",
            "is_system": false
        }
    ]
}