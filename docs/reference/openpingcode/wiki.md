知识管理
空间
用于查看和管理空间相关的基本信息。
资源地址：{GET} https://rest_api_root/v1/wiki/spaces/{space_id}

资源属性
字段	类型	描述
id	String	
空间的id。

url	String	
空间的资源地址。

identifier	String	
空间的标识。

name	String	
空间的名称。

scope_type	String	
空间的所属类型。允许值分别表示企业可见、团队可见和用户可见。

允许值: organization, user_group, user

scope_id	String	
空间的所属id。

visibility	String	
空间的可见性。

允许值: private, public

color	String	
空间的主题色。

description	String	
空间的描述。

members	Object[]	
空间的成员列表。

created_at	Number	
空间的创建时间。

created_by	Object	
空间的创建人。

updated_at	Number	
空间的更新时间。

updated_by	Object	
空间的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "示例空间描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
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
    "updated_at": 1583290400,
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
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "is_archived": 0,
    "is_deleted": 0
}
创建一个空间
企业令牌不能创建scope_type为user类型的空间。

POST
https://rest_api_root/v1/wiki/spaces
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
空间的名称（不超过32个字符）。

identifier	String	
空间的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

scope_type	String	
空间的所属类型。允许值分别表示企业可见、团队可见和用户可见。

允许值: organization, user_group, user

scope_id可选	String	
空间的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。

visibility可选	String	
空间可见性。允许值分别表示组织全部成员可见和仅空间成员可见。

允许值: public, private

description可选	String	
空间的描述。

members可选	Object[]	
空间成员的列表。

members.id	String	
企业成员或团队的id。

members.type	String	
空间成员类型。

允许值: user, user_group

请求示例：
{
    "name": "团队空间",
    "identifier": "GROUP",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "description": "团队空间所属一个团队，只能添加所属团队内的成员。",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "type": "user"
        }
    ],
}
响应示例：
{
    "id": "642fd641209b56920a6c6e5f",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5f",
    "identifier": "GROUP",
    "name": "团队空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "团队空间所属一个团队，只能添加所属团队内的成员。",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5f/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5f/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
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
部分更新一个空间
企业令牌不能更新scope_type为user类型的空间。

PATCH
https://rest_api_root/v1/wiki/spaces/{space_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
space_id	String	
空间的id。

请求参数
字段	类型	描述
name可选	String	
空间的名称（不超过32个字符）。

identifier可选	String	
空间的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
空间的描述。

请求示例：
{
    "name": "示例空间",
    "identifier": "DEMO",
    "description": "示例空间描述"
}
响应示例：
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "示例空间描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
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
    "updated_at": 1583290400,
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
获取空间列表
GET
https://rest_api_root/v1/wiki/spaces
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
keywords可选	String	
关键字。只支持name关键字搜索。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "642fd641209b56920a6c6e5e",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
            "identifier": "DEMO",
            "name": "示例空间",
            "scope_type": "user_group",
            "scope_id": "63c8fb32729dee3334d96af7",
            "visibility": "private",
            "color": "#FB7894",
            "description": "示例空间描述",
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                },
                {
                    "id": "63c8fb32729dee3334d96af7",
                    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
                    "type": "user_group",
                    "user_group": {
                        "id": "63c8fb32729dee3334d96af7",
                        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                        "name": "Open Team"
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
            "updated_at": 1583290400,
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
删除一个空间
DELETE
https://rest_api_root/v1/wiki/spaces/{space_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
space_id	String	
空间的id。

响应示例：
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "示例空间描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
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
    "updated_at": 1583290400,
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
向空间中添加一个成员
POST
https://rest_api_root/v1/wiki/spaces/{space_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
space_id	String	
空间的id。

请求参数
字段	类型	描述
member	Object	
空间的成员。

member.id	String	
企业成员或团队的id。

member.type	String	
空间成员的类型。

允许值: user, user_group

role_id可选	String	
角色的id。

请求示例：
{
    "member": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "type": "user"
    },
    "role_id": "6422711c3f12e6c1e46d40e6"
}
响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
    "space": {
        "id": "642fd641209b56920a6c6e5e",
        "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
        "identifier": "DEMO",
        "name": "示例空间",
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
获取空间中的成员列表
GET
https://rest_api_root/v1/wiki/spaces/{space_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
space_id	String	
空间的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "space": {
                "id": "642fd641209b56920a6c6e5e",
                "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
                "identifier": "DEMO",
                "name": "示例空间",
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
            "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "space": {
                "id": "642fd641209b56920a6c6e5e",
                "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
                "identifier": "DEMO",
                "name": "示例空间",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "userGroup",
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
在空间中移除一个成员
DELETE
https://rest_api_root/v1/wiki/spaces/{space_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
space_id	String	
空间的id。

member_id	String	
企业成员或团队的id。

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
    "space": {
        "id": "642fd641209b56920a6c6e5e",
        "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
        "identifier": "DEMO",
        "name": "示例空间",
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
页面
用于查看和管理页面相关的基本信息。企业令牌只能管理非个人空间下的页面。
资源地址：{GET} https://rest_api_root/v1/wiki/pages/{page_id}

资源属性
字段	类型	描述
id	String	
页面的id。

url	String	
页面的资源地址。

space	Object	
页面所属的空间。

name	String	
页面的名称。

type	String	
页面的类型。

允许值: document, group

short_id	String	
页面的短id。

html_url	String	
页面的html地址。

parent	Object	
页面的父页面。

icon	String	
页面的图标。

readings	Number	
页面的阅读量。

published_at	Number	
页面的发布时间。

published_by	Object	
页面的发布人。

participants	Object[]	
页面的关注人列表。

created_at	Number	
页面的创建时间。

created_by	Object	
页面的创建人。

updated_at	Number	
页面的更新时间。

updated_by	Object	
页面的更新人。

is_locked	Number	
是否已锁定。

允许值: 0, 1

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "示例页面",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
    },
    "icon": "file-fill",
    "readings": 10,
    "published_at": 1694065129,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
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
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1694065129,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 0,
    "is_archived": 0,
    "is_deleted": 0
}
引用数据示例：
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "name": "示例页面",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
}
创建一个页面
POST
https://rest_api_root/v1/wiki/pages
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
space_id	String	
空间的id。

name	String	
页面的名称。

parent_id可选	String	
父页面的id。

content可选	String	
页面的内容。

format_type可选	String	
页面内容的格式。content和format_type字段必须同时传递。

允许值: text, markdown, html

请求示例：
{
    "space_id": "63e1bf51760505c8795ebcc8",
    "name": "示例页面",
    "parent_id": "63e1bf51760505c8795ebcce",
    "content": "空间是记录信息和知识的页面集合，通过组织页面层级将知识系统化、结构化，便于团队沉淀经验、共享资源，实现知识增值，加快知识流动，在知识管理层面助力企业更快更好的发布产品。 PingCode 空间支持以下特性： 页面支持插入多种元素以及关联页面，满足编写需要 编辑过程自动保存草稿，无需担心内容丢失 提供丰富的模板，使用模板保持页面的一致性，让空间更加规范 使用锁定功能锁定页面最终版本 删除的页面放进回收站，随时恢复 树状页面结构，直接拖动页面编排目录，让知识管理更方便高效 通过设置成员角色来进行权限控制 通过页面评论实现成员沟通交流，形成反馈闭环  【PingCode 空间】当前处于不断迭代过程中，更多功能即将呈现，敬请期待~",
    "format_type": "text"
}
响应示例：
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "示例页面",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "c-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/c-x6NN"
    },
    "icon": "file-fill",
    "readings": 0,
    "published_at": 1675738962,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
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
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1675738962,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 0,
    "is_archived": 0,
    "is_deleted": 0
}
部分更新一个页面
PATCH
https://rest_api_root/v1/wiki/pages/{page_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
page_id	String	
页面的id。

请求参数
字段	类型	描述
name可选	String	
页面的名称。

parent_id可选	String	
父页面的id。

lock可选	Number	
是否锁定页面。

允许值: 0, 1

请求示例：
{
    "name": "示例页面updated",
    "parent_id": "63e1bf51760505c8795ebcce",
    "lock": 1
}
响应示例：
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "示例页面updated",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "c-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/c-x6NN"
    },
    "icon": "file-fill",
    "readings": 0,
    "published_at": 1675739999,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
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
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1675739999,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 1,
    "is_archived": 0,
    "is_deleted": 0
}
更新一个文档正文
PUT
https://rest_api_root/v1/wiki/pages/{page_id}/content
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
page_id	String	
页面的id。

请求参数
字段	类型	描述
content	String	
页面的内容。

format_type	String	
页面内容的格式。

允许值: text, markdown, html

请求示例：
{
    "content": "**空间是记录信息和知识的页面集合，通过组织页面层级将知识系统化、结构化，便于团队沉淀经验、共享资源，实现知识增值，加快知识流动，在知识管理层面助力企业更快更好的发布产品。** *PingCode* 空间支持以下特性： 页面支持插入多种元素以及关联页面，满足编写需要 编辑过程自动保存草稿，无需担心内容丢失 提供丰富的模板，使用模板保持页面的一致性，让空间更加规范 使用锁定功能锁定页面最终版本 删除的页面放进回收站，随时恢复 树状页面结构，直接拖动页面编排目录，让知识管理更方便高效 通过设置成员角色来进行权限控制 通过页面评论实现成员沟通交流，形成反馈闭环 **【PingCode 空间】当前处于不断迭代过程中，更多功能即将呈现，敬请期待~**",
    "format_type": "markdown"
}
响应示例：
{
    "id": "65093a8e4d4c8ca623da8fcd",
    "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/content",
    "version": {
        "id": "65093abf4d4c8ca623da8ffe",
        "url": "https://rest_api_root/v1/wiki/63e1bf51760505c8795ebccc/versions/65093abf4d4c8ca623da8ffe",
        "name": "v3",
        "order": 3
    },
    "format_type": "markdown",
    "content": "**空间是记录信息和知识的页面集合，通过组织页面层级将知识系统化、结构化，便于团队沉淀经验、共享资源，实现知识增值，加快知识流动，在知识管理层面助力企业更快更好的发布产品。** *PingCode* 空间支持以下特性： 页面支持插入多种元素以及关联页面，满足编写需要 编辑过程自动保存草稿，无需担心内容丢失 提供丰富的模板，使用模板保持页面的一致性，让空间更加规范 使用锁定功能锁定页面最终版本 删除的页面放进回收站，随时恢复 树状页面结构，直接拖动页面编排目录，让知识管理更方便高效 通过设置成员角色来进行权限控制 通过页面评论实现成员沟通交流，形成反馈闭环 **【PingCode 空间】当前处于不断迭代过程中，更多功能即将呈现，敬请期待~**"
}
获取一个文档正文
GET
https://rest_api_root/v1/wiki/pages/{page_id}/content
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
page_id	String	
页面的id。

查询参数
字段	类型	描述
format_type可选	String	
正文格式。

默认值: text

允许值: text, markdown, html

version_id可选	String	
页面版本的id。默认值为页面当前版本的id。

响应示例：
{
    "id": "65093a8e4d4c8ca623da8fcd",
    "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/content",
    "version": {
        "id": "65093abf4d4c8ca623da8ffe",
        "url": "https://rest_api_root/v1/wiki/63e1bf51760505c8795ebccc/versions/65093abf4d4c8ca623da8ffe",
        "name": "v3",
        "order": 3
    },
    "format_type": "text",
    "content": "空间是记录信息和知识的页面集合，通过组织页面层级将知识系统化、结构化，便于团队沉淀经验、共享资源，实现知识增值，加快知识流动，在知识管理层面助力企业更快更好的发布产品。 PingCode 空间支持以下特性： 页面支持插入多种元素以及关联页面，满足编写需要 编辑过程自动保存草稿，无需担心内容丢失 提供丰富的模板，使用模板保持页面的一致性，让空间更加规范 使用锁定功能锁定页面最终版本 删除的页面放进回收站，随时恢复 树状页面结构，直接拖动页面编排目录，让知识管理更方便高效 通过设置成员角色来进行权限控制 通过页面评论实现成员沟通交流，形成反馈闭环  【PingCode 空间】当前处于不断迭代过程中，更多功能即将呈现，敬请期待~"
}
获取一个页面的版本列表
GET
https://rest_api_root/v1/wiki/pages/{page_id}/versions
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
page_id	String	
页面的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "65093abf4d4c8ca623da8ffe",
            "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/versions/65093abf4d4c8ca623da8ffe",
            "page": {
                "id": "65093a8e4d4c8ca623da8fcd",
                "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd",
                "name": "主页",
                "type": "document",
                "short_id": "AAx6NN",
                "html_url": "https://yctech.pingcode.com/wiki/pages/AAx6NN"
            },
            "name": "v1",
            "order": 1,
            "status": "published",
            "created_at": 1695103679,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
恢复一个页面到指定版本
POST
https://rest_api_root/v1/wiki/pages/{page_id}/versions/{version_id}/restore
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
page_id	String	
页面的id。

version_id	String	
页面版本的id。

响应示例：
{
    "id": "65093abf4d4c8ca623da8fff",
    "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/versions/65093abf4d4c8ca623da8fff",
    "page": {
        "id": "65093a8e4d4c8ca623da8fcd",
        "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd",
        "name": "主页",
        "type": "document",
        "short_id": "5-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
    },
    "name": "v2恢复自v1",
    "order": 2,
    "status": "published",
    "created_at": 1695103832,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
获取页面列表
GET
https://rest_api_root/v1/wiki/pages?space_id={space_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
space_id	String	
空间的id。

parent_id可选	String	
父页面的id。

created_between可选	String	
创建时间介于的时间范围，通过','分割起始时间。

updated_between可选	String	
更新时间介于的时间范围，通过','分割起始时间。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "63e1bf51760505c8795ebccc",
            "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
            "space": {
                "id": "63e1bf51760505c8795ebcc8",
                "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
                "name": "示例空间",
                "identifier": "DEMO",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "示例页面",
            "type": "document",
            "short_id": "5-x6NN",
            "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
            "parent": {
                "id": "63e1bf51760505c8795ebcce",
                "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
                "name": "模板",
                "type": "document",
                "short_id": "5-x6NN",
                "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
            },
            "icon": "file-fill",
            "readings": 10,
            "published_at": 1694065129,
            "published_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
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
            "created_at": 1675738962,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1694065129,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_locked": 0,
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "63e1bf51760505c8795ebcce",
            "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
            "space": {
                "id": "63e1bf51760505c8795ebcc8",
                "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
                "name": "示例空间",
                "identifier": "DEMO",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "模板",
            "type": "document",
            "short_id": "5-x6NN",
            "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
            "parent": null,
            "icon": "file-fill",
            "readings": 0,
            "published_at": 1694065129,
            "published_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebcce",
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
            "created_at": 1675738962,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1694065129,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_locked": 0,
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
删除一个页面
DELETE
https://rest_api_root/v1/wiki/pages/{page_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
page_id	String	
页面的id。

响应示例：
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "name": "示例页面updated",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "c-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/c-x6NN"
    },
    "icon": "file-fill",
    "readings": 0,
    "published_at": 1675739999,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
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
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1675739999,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 0,
    "is_archived": 0,
    "is_deleted": 1
}