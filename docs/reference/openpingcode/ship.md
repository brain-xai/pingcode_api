产品管理
产品
用于查看和管理产品相关的基本信息。
资源地址：{GET} /v1/ship/products/{product_id}

资源属性
字段	类型	描述
id	String	
产品的id。

url	String	
产品的资源地址。

identifier	String	
产品的标识。

name	String	
产品的名称。

scope_type	String	
产品的所属类型。

允许值: organization, user_group

scope_id	String	
产品的所属id。

visibility	String	
产品的可见性。

允许值: private, public

color	String	
产品的主题色。

description	String	
产品的描述。

members	Object[]	
产品的成员列表。

created_at	Number	
产品的创建时间。

created_by	Object	
产品的创建人。

updated_at	Number	
产品的更新时间。

updated_by	Object	
产品的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "scope_type": "user_group",
    "scope_id": "6422711c3f12e6c1e46d40e9",
    "visibility": "private",
    "color": "#FA8888",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
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
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/63c8fb32729dee3334d96af7",
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
引用数据示例：
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "is_archived": 0,
    "is_deleted": 0
}
创建一个产品
POST
https://rest_api_root/v1/ship/products
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
产品的名称（不超过32个字符）。

identifier	String	
产品的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
产品的描述。

members可选	Object[]	
产品成员的列表。

members.id	String	
企业成员或团队的id。

members.type	String	
产品成员的类型。

允许值: user, user_group

请求示例：
{
    "name": "示例产品",
    "identifier": "SLC",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "type": "user"
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "type": "user_group"
        }
    ]
}
响应示例：
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "visibility": "private",
    "scope_type": "user_group",
    "scope_id": "6422711c3f12e6c1e46d40e9",
    "color": "#FA8888",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
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
部分更新一个产品
PATCH
https://rest_api_root/v1/ship/products/{product_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

请求参数
字段	类型	描述
name可选	String	
产品的名称（不超过32个字符）。

identifier可选	String	
产品的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
产品的描述。

请求示例：
{
    "name": "示例产品",
    "identifier": "SLC",
    "description": "示例产品描述"
}
响应示例：
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "scope_type": "user_group",
    "scope_id": "6422711c3f12e6c1e46d40e9",
    "visibility": "private",
    "color": "#FA8888",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
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
获取产品列表
GET
https://rest_api_root/v1/ship/products
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
            "identifier": "SLC",
            "name": "示例产品",
            "visibility": "private",
            "scope_type": "user_group",
            "scope_id": "6422711c3f12e6c1e46d40e9",
            "color": "#FA8888",
            "description": "示例产品描述",
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
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
向产品中添加一个成员
POST
https://rest_api_root/v1/ship/products/{product_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

请求参数
字段	类型	描述
member	Object	
产品的成员。

member.id	String	
企业成员或团队的id。

member.type	String	
项目成员的类型。

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
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "name": "示例产品",
        "identifier": "SLC",
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
获取产品中的成员列表
GET
https://rest_api_root/v1/ship/products/{product_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "name": "示例产品",
                "identifier": "SLC",
                "is_archived": 0,
                "is_deleted": 0
            },
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
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/63c8fb32729dee3334d96af7",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "name": "示例产品",
                "identifier": "SLC",
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
在产品中移除一个成员
DELETE
https://rest_api_root/v1/ship/products/{product_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

member_id	String	
企业成员或团队的id。

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "name": "示例产品",
        "identifier": "SLC",
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
向产品中添加一个需求模块
POST
https://rest_api_root/v1/ship/products/{product_id}/suites
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

请求参数
字段	类型	描述
name	String	
模块名称。需求模块为树形结构，同一层次下的名称不能重复。

type	String	
模块类型。允许值分别表示子产品和模块。

允许值: product, module

parent_id可选	String	
父模块的id。

请求示例：
{
    "name": "技术支持确认",
    "type": "module",
    "parent_id": "63bb744414bd13c9def24ce3"
}
响应示例：
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "技术支持确认",
    "type": "module",
    "parent": {
        "id": "63bb744414bd13c9def24ce3",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63bb744414bd13c9def24ce3",
        "name": "父级产品模块",
        "type": "module"
    }
}
获取产品中的需求模块列表
GET
https://rest_api_root/v1/ship/products/{product_id}/suites
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49f0",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63eca881a0a13a3efc8d49f0",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "技术支持确认",
            "type": "module",
            "parent": {
                "id": "63bb744414bd13c9def24ce3",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63bb744414bd13c9def24ce3",
                "name": "父级产品模块",
                "type": "module"
            }
        }
    ]
}
在产品中移除一个需求模块
请注意，删除一个模块会自动删除其所有的子模块。

DELETE
https://rest_api_root/v1/ship/products/{product_id}/suites/{suite_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

suite_id	String	
模块id。

响应示例：
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "技术支持确认",
    "type": "module",
    "parent": {
        "id": "63bb744414bd13c9def24ce3",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63bb744414bd13c9def24ce3",
        "name": "父级产品模块",
        "type": "module"
    }
}
获取产品中的需求排期列表
GET
https://rest_api_root/v1/ship/products/{product_id}/plans
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744414bd13c9def24ce4",
            "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/plans/63bb744414bd13c9def24ce4",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "账号管理",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1672704000,
            "end_at": 1672963199
        }
    ]
}
获取产品中的工单渠道列表
GET
https://rest_api_root/v1/ship/products/{product_id}/channels
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49ed",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/channels/63eca881a0a13a3efc8d49ed",
            "name": "客户反馈Web渠道",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "description": "收集客户反馈意见的Web渠道"
        }
    ]
}
获取产品中的工单类型列表
GET
https://rest_api_root/v1/ship/products/{product_id}/ticket_types
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/ticket_types/63bb744214bd13c9def24ca9",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "name": "自动化",
                "identifier": "FLOW",
                "is_archived": 0,
                "is_deleted": 0
            },
            "ticket_type": {
                "id": "63bb744214bd13c9def24ca9",
                "url": "https://rest_api_root/v1/ship/ticket_types/63bb744214bd13c9def24ca9",
                "name": "需求"
            }
        }
    ]
}
获取产品中的标签列表
GET
https://rest_api_root/v1/ship/products/{product_id}/tags
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49f0",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "技术支持确认",
            "color": "#56ABFB"
        }
    ]
}
向产品中添加一个标签
POST
https://rest_api_root/v1/ship/products/{product_id}/tags
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

请求参数
字段	类型	描述
name	String	
标签的名称。在一个产品中这个名称是唯一的。

请求示例：
{
    "name": "标签-1"
}
响应示例：
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "标签-1",
    "color": "#56ABFB"
}
在产品中移除一个标签
DELETE
https://rest_api_root/v1/ship/products/{product_id}/tags/{tag_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

tag_id	String	
标签的id。

响应示例：
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "技术支持确认",
    "color": "#56ABFB"
}
客户
用于查看和管理客户相关的基本信息。
资源地址：{GET} https://rest_api_root/v1/ship/products/{product_id}/customers/{customer_id}

资源属性
字段	类型	描述
id	String	
客户的id。

url	String	
客户的资源地址。

product	Object	
客户的所属产品。

name	String	
客户的名称。

assignee	Object	
客户的负责人。

scale	Number	
客户的规模。

description	String	
客户的描述。

created_at	Number	
客户的创建时间。

created_by	Object	
客户的创建人。

updated_at	Number	
客户的更新时间。

updated_by	Object	
客户的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "上海XX新零售有限公司",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scale": 200,
    "description": "上海XX新零售有限公司的描述",
    "created_at": 1692240286,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1692240286,
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
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "name": "上海XX新零售有限公司"
}
创建一个客户
POST
https://rest_api_root/v1/ship/products/{product_id}/customers
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

请求参数
字段	类型	描述
name	String	
客户的名称。

assignee_id可选	String	
客户的负责人id。

scale可选	Number	
客户的规模。

description可选	String	
客户的描述。

请求示例：
{
    "name": "上海XX新零售有限公司",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "scale": 200,
    "description": "上海XX新零售有限公司的描述"
}
响应示例：
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "上海XX新零售有限公司",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scale": 200,
    "description": "上海XX新零售有限公司的描述",
    "created_at": 1692240286,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1692240286,
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
部分更新一个客户
PATCH
https://rest_api_root/v1/ship/products/{product_id}/customers/{customer_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

customer_id	String	
客户的id。

请求参数
字段	类型	描述
name可选	String	
客户的名称。

assignee_id可选	String	
客户的负责人id。

scale可选	Number	
客户的规模。

description可选	String	
客户的描述。

请求示例：
{
    "name": "上海XX新零售有限公司",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "scale": 200,
    "description": "上海XX新零售有限公司的描述"
}
响应示例：
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "上海XX新零售有限公司",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scale": 200,
    "description": "上海XX新零售有限公司的描述",
    "created_at": 1692240286,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1692241286,
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
获取客户列表
GET
https://rest_api_root/v1/ship/products/{product_id}/customers
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64dd899e3f6383ba72ec2a0d",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "上海XX新零售有限公司",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "scale": 200,
            "description": "上海XX新零售有限公司的描述",
            "created_at": 1692240286,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1692240286,
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
外部用户
用于查看和管理外部用户相关的基本信息。
资源地址：{GET} /v1/ship/products/{product_id}/users/{user_id}

资源属性
字段	类型	描述
id	String	
外部用户的id。

url	String	
外部用户的资源地址。

name	String	
外部用户的名称。

display_name	String	
外部用户的显示名。

avatar	String	
外部用户的头像地址。

email	String	
外部用户的邮箱地址。

mobile	String	
外部用户的手机号。

product	Object	
外部用户所属的产品。

customer	Object	
外部用户所属的客户。

全量数据示例：
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
引用数据示例：
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png"
}
创建一个外部用户
POST
https://rest_api_root/v1/ship/products/{product_id}/users
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

请求参数
字段	类型	描述
name	String	
外部用户的名称。

email可选	String	
外部用户的邮箱地址，邮箱地址和手机号其中一个必填。

mobile可选	String	
外部用户的手机号，邮箱地址和手机号其中一个必填，如果同时存在，以手机号为准。

customer_id可选	String	
外部用户所属客户的id。

请求示例：
{
    "name": "jack",
    "email": "jack@email.com",
    "customer_id": "64dd899e3f6383ba72ec2a01"
}
响应示例：
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
部分更新一个外部用户
PATCH
https://rest_api_root/v1/ship/products/{product_id}/users/{user_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

user_id	String	
外部用户的id。

请求示例：
{
    "customer_id": "64dd899e3f6383ba72ec2a01"
}
Success 200
字段	类型	描述
customer_id可选	String	
外部用户所属客户的id。

响应示例：
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
获取外部用户列表
GET
https://rest_api_root/v1/ship/products/{product_id}/users
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64a2b61c3a12e6c2e46d41e9",
            "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
            "name": "jack",
            "display_name": "Jack",
            "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
            "email": "jack@email.com",
            "mobile": null,
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "customer": {
                "id": "64dd899e3f6383ba72ec2a01",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
                "name": "深圳XX新零售有限公司"
            }
        }
    ]
}
删除一个外部用户
DELETE
https://rest_api_root/v1/ship/products/{product_id}/users/{user_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
product_id	String	
产品的id。

user_id	String	
外部用户的id。

响应示例：
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
工单
用于查看和管理工单相关的基本信息。
资源地址：{GET} /v1/ship/tickets/{ticket_id}

资源属性
字段	类型	描述
id	String	
工单的id。

url	String	
工单的资源地址。

product	Object	
工单的所属产品。

identifier	String	
工单的标识。

title	String	
工单的标题。

short_id	String	
工单的短id。

html_url	String	
工单的html地址。

assignee	Object	
工单的负责人。

state	Object	
工单的状态。

type	Object	
工单的类型。

customer	Object	
工单的客户。

solution	Object	
工单的解决方案。

priority	Object	
工单的优先级。

channel	Object/String	
工单的渠道。外部渠道提交的工单的渠道是Object类型，内部工单的渠道是 internal 字符串。

description	String	
工单的描述。

properties	Object	
工单的自定义属性。

properties.prop_a	Object	
工单的自定义属性prop_a。

properties.prop_b	Object	
工单的自定义属性prop_b。

estimated_at	Object	
工单的预计时间。

estimated_at.from	Number	
预计时间的开始时间。

estimated_at.to	Number	
预计时间的结束时间。

estimated_at.granularity	String	
预计时间的周期单位。

允许值: year, quarter, month, day

tags	Object[]	
工单的标签列表。

participants	Object[]	
工单的关注人列表。

public_image_token	String	
工单描述和自定义多行文本类型属性里获取图片资源token。获取一个工单和获取工单列表参数include_public_image_token值有效时返回。

submitted_at	Number	
工单的提交时间。

submitted_by	Object	
工单的提交人。

completed_at	Number	
工单的完成时间。

completed_by	Object	
工单的完成人。

created_at	Number	
工单的创建时间。

created_by	Object	
工单的创建人。

updated_at	Number	
工单的更新时间。

updated_by	Object	
工单的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63eca880a0a13a3efc8d49d9",
        "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
        "name": "待处理",
        "type": "pending"
    },
    "estimated_at": {
        "from": 1701619200,
        "to": 1702742399,
        "granularity": "day"
    },
    "type": {
        "id": "63eca880a0a13a3efc8d49e0",
        "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
        "name": "需求"
    },
    "customer": {
        "id": "63eca881a0a13a3efc8d49fc",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
        "name": "北京XX科技有限公司"
    },
    "solution": {
        "id": "62f217ae16e3661a20124330",
        "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
        "name": "进入需求池"
    },
    "priority": {
        "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
        "id": "5cb9466afda1ce4ca0090004",
        "name": "P1"
    },
    "channel": {
        "id": "64550d9ec696b249b5fc607d",
        "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
        "name": "channel-1"
    },
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "tags": [
        {
            "id": "63eca881a0a13a3efc8d49f1",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/tags/63eca881a0a13a3efc8d49f1",
            "name": "已确认"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
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
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "public_image_token": "N96GlJ4AMQoBCw9pZQ2EMl-AprLN_V_DYlghupBNkJA",
    "submitted_at": 1676454024,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676454024,
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
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
}
创建一个工单
POST
https://rest_api_root/v1/ship/tickets
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
product_id	String	
工单的产品id。

title	String	
工单的标题，最大长度为255。

type_id	String	
工单的类型id，您可以在 “获取工单类型列表” API获取。

description可选	String	
工单的描述。

submitter_id可选	String	
工单的提交人id，企业授权时，该值有效；个人鉴权时，指定无效。

customer_id可选	String	
工单的客户id，您可以在 “获取产品客户列表” API获取。

channel_id可选	String	
工单的渠道id，您可以在 “获取渠道列表” API获取。

assignee_id可选	String	
工单的负责人id。

priority_id可选	String	
工单的优先级id，您可以在 “获取工单优先级列表” API获取。

properties可选	Object	
工单的自定义属性。

properties.prop_a可选	Object	
工单的自定义属性prop_a。

properties.prop_b可选	Object	
工单的自定义属性prop_b。

请求示例：
{
    "product_id": "6422711c3f12e6c1e46d40e9",
    "title": "希望新增支持第三方账号注册",
    "type_id": "63eca880a0a13a3efc8d49e0",
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "submitter_id": "a0417f68e846aae315c85d24643678a9",
    "customer_id": "63eca881a0a13a3efc8d49fc",
    "channel_id": "64550d9ec696b249b5fc607d",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "priority_id": "5cb9466afda1ce4ca0090004",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63eca880a0a13a3efc8d49d9",
        "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
        "name": "待处理",
        "type": "pending"
    },
    "type": {
        "id": "63eca880a0a13a3efc8d49e0",
        "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
        "name": "需求"
    },
    "customer": {
        "id": "63eca881a0a13a3efc8d49fc",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
        "name": "北京XX科技有限公司"
    },
    "solution": {
        "id": "62f217ae16e3661a20124330",
        "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
        "name": "进入需求池"
    },
    "priority": {
        "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
        "id": "5cb9466afda1ce4ca0090004",
        "name": "P1"
    },
    "channel": {
        "id": "64550d9ec696b249b5fc607d",
        "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
        "name": "channel-1"
    },
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_at": null,
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
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
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "submitted_at": 1676454024,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676454024,
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
部分更新一个工单
PATCH
https://rest_api_root/v1/ship/tickets/{ticket_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
ticket_id	String	
工单id。

请求参数
字段	类型	描述
title可选	String	
工单的标题，最大长度为255。

description可选	String	
工单的描述。

type_id可选	String	
工单的类型id，您可以在 “获取工单类型列表” API获取。

state_id可选	String	
工单的状态id，您可以在 “获取工单状态列表” API获取。

assignee_id可选	String	
工单的负责人id。

submitter_id可选	String	
工单的提交人id，企业授权时，该值有效；个人鉴权时，指定无效。

solution_id可选	String	
工单的解决方案id，您可以在 “获取工单解决方案列表” API获取。

priority_id可选	String	
工单的优先级id，您可以在 “获取工单优先级列表” API获取。

customer_id可选	String	
工单的客户id，您可以在 “获取产品客户列表” API获取。

properties可选	Object	
工单的自定义属性。

properties.prop_a可选	Object	
工单的自定义属性prop_a。

properties.prop_b可选	Object	
工单的自定义属性prop_b。

请求示例：
{
    "title": "希望新增支持第三方账号注册",
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "type_id": "63eca880a0a13a3efc8d49e0",
    "state_id": "63eca880a0a13a3efc8d49e0",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "submitter_id": "a0417f68e846aae315c85d24643678a9",
    "solution_id": "62f217ae16e3661a20124330",
    "priority_id": "5cb9466afda1ce4ca0090004",
    "customer_id": "63eca881a0a13a3efc8d49fc",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63eca880a0a13a3efc8d49d9",
        "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
        "name": "待处理",
        "type": "pending"
    },
    "estimated_at": {
        "from": 1701619200,
        "to": 1702742399,
        "granularity": "day"
    },
    "type": {
        "id": "63eca880a0a13a3efc8d49e0",
        "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
        "name": "需求"
    },
    "customer": {
        "id": "63eca881a0a13a3efc8d49fc",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
        "name": "北京XX科技有限公司"
    },
    "solution": {
        "id": "62f217ae16e3661a20124330",
        "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
        "name": "进入需求池"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090004",
        "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
        "name": "P1"
    },
    "channel": {
        "id": "64550d9ec696b249b5fc607d",
        "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
        "name": "channel-1"
    },
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
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
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "submitted_at": 1676454024,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676455024,
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
获取工单列表
GET
https://rest_api_root/v1/ship/tickets
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id可选	String	
产品的id。

type_id可选	String	
工单类型id。

state_id可选	String	
工单状态id。

priority_id可选	String	
工单优先级id。

created_between可选	String	
创建时间介于的时间范围，通过','分割起始时间。

updated_between可选	String	
更新时间介于的时间范围，通过','分割起始时间。

keywords可选	String	
关键字。支持工单编号和工单标题。

include_public_image_token可选	String	
包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.prop_b。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca888a0a13a3efc8d4a43",
            "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SLC-T1",
            "title": "希望新增支持第三方账号注册",
            "short_id": "pdMUgQzZ",
            "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "state": {
                "id": "63eca880a0a13a3efc8d49d9",
                "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
                "name": "待处理",
                "type": "pending"
            },
            "estimated_at": {
                "from": 1701619200,
                "to": 1702742399,
                "granularity": "day"
            },
            "type": {
                "id": "63eca880a0a13a3efc8d49e0",
                "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
                "name": "需求"
            },
            "customer": {
                "id": "63eca881a0a13a3efc8d49fc",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
                "name": "北京XX科技有限公司"
            },
            "solution": {
                "id": "62f217ae16e3661a20124330",
                "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
                "name": "进入需求池"
            },
            "priority": {
                "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
                "id": "5cb9466afda1ce4ca0090004",
                "name": "P1"
            },
            "channel": {
                "id": "64550d9ec696b249b5fc607d",
                "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
                "name": "channel-1"
            },
            "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "tags": [
                {
                    "id": "63eca881a0a13a3efc8d49f1",
                    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/tags/63eca881a0a13a3efc8d49f1",
                    "name": "已确认"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
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
                    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
                    "type": "user_group",
                    "user_group": {
                        "id": "63c8fb32729dee3334d96af7",
                        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                        "name": "Open Team"
                    }
                }
            ],
            "public_image_token": "N96GlJ4AMQoBCw9pZQ2EMl-AprLN_V_DYlghupBNkJA",
            "submitted_at": 1676454024,
            "submitted_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "completed_at": 1689579131,
            "completed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "created_at": 1676454024,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1676454024,
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
获取工单类型列表
GET
https://rest_api_root/v1/ship/ticket/types?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca9",
            "url": "https://rest_api_root/v1/ship/ticket_types/63bb744214bd13c9def24ca9",
            "name": "需求"
        }
    ]
}
获取工单状态列表
GET
https://rest_api_root/v1/ship/ticket/states?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40f2",
            "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
            "name": "处理中",
            "type": "pending"
        }
    ]
}
获取工单属性列表
GET
https://rest_api_root/v1/ship/ticket/properties?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
            "name": "解决方案",
            "type": "select",
            "options": [
                {
                    "_id": "6422711c3f12e6c1e46d40e9",
                    "text": "进入需求池"
                }
            ]
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/ticket_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null
        }
    ]
}
获取工单渠道列表
GET
https://rest_api_root/v1/ship/ticket/channels?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49ed",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/channels/63eca881a0a13a3efc8d49ed",
            "name": "客户反馈Web渠道"
        }
    ]
}
获取工单优先级列表
GET
https://rest_api_root/v1/ship/ticket/priorities?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cb9466afda1ce4ca0090005",
            "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090005",
            "name": "P0"
        }
    ]
}
获取工单解决方案列表
GET
https://rest_api_root/v1/ship/ticket/solutions?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e9",
            "url": "https://rest_api_root/v1/ship/ticket_solutions/6422711c3f12e6c1e46d40e9",
            "name": "进入需求池"
        }
    ]
}
获取工单标签列表
GET
https://rest_api_root/v1/ship/ticket/tags?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49f0",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
            "name": "技术支持确认"
        }
    ]
}
获取工单流转记录列表
GET
https://rest_api_root/v1/ship/tickets/{ticket_id}/transition_histories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
ticket_id	String	
工单的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "64c3676c983bb9481ee1eea5",
            "url": "https://rest_api_root/v1/ship/tickets/5edca524cad2fa1125cb0630/transition_histories/64c3676c983bb9481ee1eea5",
            "ticket": {
                "id": "63eca888a0a13a3efc8d4a43",
                "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
                "identifier": "SLC-T1",
                "title": "希望新增支持第三方账号注册",
                "short_id": "pdMUgQzZ",
                "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
            },
            "from_state": null,
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/ticket_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "658bdb79e5839f556562accf",
            "url": "https://rest_api_root/v1/ship/tickets/5edca524cad2fa1125cb0630/transition_histories/658bdb79e5839f556562accf",
            "ticket": {
                "id": "63eca888a0a13a3efc8d4a43",
                "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
                "identifier": "SLC-T1",
                "title": "希望新增支持第三方账号注册",
                "short_id": "pdMUgQzZ",
                "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
            },
            "from_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/ticket_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b2b",
                "url": "https://rest_api_root/v1/ship/ticket_states/63e1bf51898a0be5a2d21b2b",
                "name": "处理中",
                "type": "in_progress"
            },
            "created_at": 1674528614,
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
需求
用于查看和管理需求相关的基本信息。
资源地址：{GET} /v1/ship/ideas/{idea_id}

资源属性
字段	类型	描述
id	String	
需求的id。

url	String	
需求的资源地址。

product	Object	
需求的所属产品。

identifier	String	
需求的标识。

title	String	
需求的标题。

short_id	String	
需求的短id。

html_url	String	
需求的html地址。

assignee	Object	
需求的负责人。

state	Object	
需求的状态。

priority	Object	
需求的优先级。

plan	Object	
需求的计划。

suite	Object	
需求的模块。

plan_at	Object	
需求的计划时间范围。

plan_at.from	Number	
需求的计划开始时间。

plan_at.to	Number	
需求的计划结束时间。

plan_at.granularity	String	
需求的计划时间周期单位。

允许值: year, quarter, month, day

real_at	Object	
需求的实际时间范围。

real_at.from	Number	
需求的实际开始时间。

real_at.to	Number	
需求的实际结束时间。

real_at.granularity	String	
需求的计划时间周期单位。

允许值: year, quarter, month, day

score	Number	
需求的分数。

progress	Number	
需求的进度。

description	String	
需求的描述。

properties	Object	
需求的自定义属性。

properties.prop_a	Object	
需求的自定义属性prop_a。

properties.prop_b	Object	
需求的自定义属性prop_b。

participants	Object[]	
需求的关注人列表。

public_image_token	String	
需求描述和自定义多行文本类型属性里获取图片资源token。获取一个需求和获取需求列表参数include_public_image_token值有效时返回。

completed_at	Number	
需求的完成时间。

completed_by	Object	
需求的完成人。

created_at	Number	
需求的创建时间。

created_by	Object	
需求的创建人。

updated_at	Number	
需求的更新时间。

updated_by	Object	
需求的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63e1bf51898a0be5a2d21b2a",
        "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
        "name": "待评审",
        "type": "pending"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090005",
        "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
        "name": "P0"
    },
    "plan": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/plans/63bb744414bd13c9def24ce4",
        "name": "账号管理"
    },
    "suite": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
        "name": "需求模块",
        "type": "module"
    },
    "plan_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "real_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "score": 0,
    "progress": 0.6,
    "description": "这是一段描述",
    "properties": {
        "backlog_from": "5cb7e6e2fda1ce4ca0000001",
        "backlog_type": "5cb7e763fda1ce4ca0010002"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
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
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "public_image_token": "-fkvANQ2dcVECK6Xg45L3kG8VCbOTK8NrNckGkxljQD",
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1689573131,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1689579131,
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
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
}
创建一个需求
POST
https://rest_api_root/v1/ship/ideas
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
product_id	String	
需求的产品id。

title	String	
需求的标题，最大长度为255。

assignee_id可选	String	
需求负责人的id。

description可选	String	
需求的描述。

suite_id可选	String	
需求的产品模块id。

properties可选	Object	
需求属性的键值对集合。要注意的是，当前产品的需求属性视图需要包含这些需求属性，例如需求属性视图中包含prop_a和prop_b。

properties.prop_a可选	Object	
需求项属性prop_a。

properties.prop_b可选	Object	
需求项属性prop_b。

请求示例：
{
    "product_id": "6422711c3f12e6c1e46d40e9",
    "title": "示例需求",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "description": "这是一段描述",
    "suite_id": "63bb744414bd13c9def24ce4",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63e1bf51898a0be5a2d21b2a",
        "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
        "name": "待评审",
        "type": "pending"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090001",
        "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090001",
        "name": "P4"
    },
    "plan": null,
    "suite": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
        "name": "需求模块",
        "type": "module"
     },
    "plan_at": null,
    "real_at": null,
    "score": 0,
    "progress": 0,
    "description": "这是一段描述",
    "properties": {
        "backlog_from": "5cb7e6e2fda1ce4ca0000001",
        "backlog_type": "5cb7e763fda1ce4ca0010002",
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
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
    "completed_at": null,
    "completed_by": null,
    "created_at": 1689573131,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1689573131,
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
部分更新一个需求
PATCH
https://rest_api_root/v1/ship/ideas/{idea_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
idea_id	String	
需求id。

请求参数
字段	类型	描述
title可选	String	
需求的标题，最大长度为255。

description可选	String	
需求的描述。

state_id可选	String	
需求状态的id，您可以在 获取需求状态列表 API获取。

priority_id可选	String	
需求优先级的id，您可以在 获取需求优先级列表 API获取。

assignee_id可选	String	
需求负责人的id。

progress可选	Number	
需求的进度，取值范围为：0到1的两位小数。

plan_at可选	Object	
需求的计划时间。plan_at是整体更新的，其中包含from、to、granularity三个属性，均为必填。

plan_at.from	Number	
需求的计划开始时间。为秒级时间戳。

plan_at.to	Number	
需求的计划结束时间。为秒级时间戳。

plan_at.granularity	String	
需求的计划时间周期单位。

允许值: year, quarter, month, day

real_at可选	Object	
需求的实际时间，参数格式同plan_at。

plan_id可选	String	
产品排期的id，您可以在 获取产品排期列表 API获取。

suite_id可选	String	
产品模块的id，您可以在 获取产品模块列表 API获取。

properties可选	Object	
需求的自定义属性。

properties.prop_a可选	Object	
需求的自定义属性prop_a。

properties.prop_b可选	Object	
需求的自定义属性prop_b。

请求示例：
{
    "title": "示例需求",
    "description": "这是一段描述",
    "state_id": "63e1bf51898a0be5a2d21b2a",
    "priority_id": "5cb9466afda1ce4ca0090005",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "progress": 0.88,
    "plan_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "real_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "plan_id": "63bb744414bd13c9def24ce4",
    "suite_id": "63bb744414bd13c9def24ce4",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63e1bf51898a0be5a2d21b2a",
        "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
        "name": "待评审",
        "type": "pending"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090005",
        "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
        "name": "P0"
    },
    "plan": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/plans/63bb744414bd13c9def24ce4",
        "name": "账号管理"
    },
    "suite": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
        "name": "需求模块",
        "type": "module"
    },
    "plan_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "real_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "score": 0,
    "progress": 0.88,
    "description": "这是一段描述",
    "properties": {
        "backlog_from": "5cb7e6e2fda1ce4ca0000001",
        "backlog_type": "5cb7e763fda1ce4ca0010002",
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
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
    "completed_at": 1689578888,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1689573131,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1689578888,
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
获取需求列表
GET
https://rest_api_root/v1/ship/ideas
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id可选	String	
产品的id。

state_id可选	String	
需求状态id。

priority_id可选	String	
需求优先级id。

created_between可选	String	
创建时间介于的时间范围，通过','分割起始时间。

updated_between可选	String	
更新时间介于的时间范围，通过','分割起始时间。

keywords可选	String	
搜索关键字。支持需求编号和需求标题。

include_public_image_token可选	String	
包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.xxx。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64b4d70ba368e6594360ea24",
            "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SLC-1",
            "title": "示例需求",
            "short_id": "Ogf1EYey",
            "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "state": {
                "id": "63e1bf51898a0be5a2d21b2a",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
                "name": "待评审",
                "type": "pending"
            },
            "priority": {
                "id": "5cb9466afda1ce4ca0090005",
                "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
                "name": "P0"
            },
            "plan": {
                "id": "63bb744414bd13c9def24ce4",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/plans/63bb744414bd13c9def24ce4",
                "name": "账号管理"
            },
            "suite": {
                "id": "63bb744414bd13c9def24ce4",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
                "name": "需求模块",
                "type": "module"
            },
            "plan_at": {
                "from": 1690732800,
                "to": 1691337599,
                "granularity": "day"
            },
            "real_at": {
                "from": 1690732800,
                "to": 1691337599,
                "granularity": "day"
            },
            "score": 0,
            "progress": 0.6,
            "description": "这是一段描述",
            "properties": {
                "backlog_from": "5cb7e6e2fda1ce4ca0000001",
                "backlog_type": "5cb7e763fda1ce4ca0010002"
            },
            "public_image_token": "-fkvANQ2dcVECK6Xg45L3kG8VCbOTK8NrNckGkxljQd",
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
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
                    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
                    "type": "user_group",
                    "user_group": {
                        "id": "63c8fb32729dee3334d96af7",
                        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                        "name": "Open Team"
                    }
                }
            ],
            "completed_at": 1689573131,
            "completed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "created_at": 1689573131,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1689579131,
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
获取需求状态列表
GET
https://rest_api_root/v1/ship/idea/states?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63e1bf51898a0be5a2d21b2a",
            "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
            "name": "待评审",
            "type": "pending"
        }
    ]
}
获取需求属性列表
GET
https://rest_api_root/v1/ship/idea/properties?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "backlog_type",
            "url": "https://rest_api_root/v1/ship/idea_properties/backlog_type",
            "name": "需求类型",
            "type": "select",
            "options": [
                {
                    "_id": "5cb7e763fda1ce4ca0010002",
                    "text": "功能需求"
                },
                {
                    "_id": "5cb7e763fda1ce4ca0010004",
                    "text": "体验优化"
                }
            ]
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/idea_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null
        }
    ]
}
获取需求模块列表
GET
https://rest_api_root/v1/ship/idea/suites?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744414bd13c9def24ce4",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
            "name": "需求模块",
            "type": "module"
        }
    ]
}
获取需求排期列表
GET
https://rest_api_root/v1/ship/idea/plans?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744414bd13c9def24ce4",
            "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/plans/63bb744414bd13c9def24ce4",
            "name": "账号管理"
        }
    ]
}
获取需求优先级列表
GET
https://rest_api_root/v1/ship/idea/priorities?product_id={product_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
product_id	String	
产品的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cb9466afda1ce4ca0090005",
            "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
            "name": "P0"
        }
    ]
}
获取需求流转记录列表
GET
https://rest_api_root/v1/ship/ideas/{idea_id}/transition_histories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
idea_id	String	
需求的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "64c3676c983bb9481ee1eea5",
            "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24/transition_histories/64c3676c983bb9481ee1eea5",
            "idea": {
                "id": "64b4d70ba368e6594360ea24",
                "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
                "identifier": "SLC-1",
                "title": "示例需求",
                "short_id": "Ogf1EYey",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
            },
            "from_state": null,
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "658bdb79e5839f556562accf",
            "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24/transition_histories/658bdb79e5839f556562accf",
            "idea": {
                "id": "64b4d70ba368e6594360ea24",
                "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
                "identifier": "SLC-1",
                "title": "示例需求",
                "short_id": "Ogf1EYey",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
            },
            "from_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b2b",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2b",
                "name": "处理中",
                "type": "in_progress"
            },
            "created_at": 1674528614,
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
产品配置中心
用于查看和管理产品相关的配置信息。

工单配置
用于查看和管理工单相关的配置信息。

获取全部工单类型列表
GET
https://rest_api_root/v1/ship/ticket_types
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca9",
            "url": "https://rest_api_root/v1/ship/ticket_types/63bb744214bd13c9def24ca9",
            "name": "需求",
            "is_system": true
        }
    ]
}
创建一个工单状态
POST
https://rest_api_root/v1/ship/ticket_states
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
工单状态的名称，在一个企业中这个名称是唯一的。

type	String	
工单状态的类型。

允许值: pending, in_progress, completed, closed

请求示例：
{
    "name": "处理中",
    "type": "pending"
}
响应示例：
{
    "id": "6422711c3f12e6c1e46d40f2",
    "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
    "name": "处理中",
    "type": "pending",
    "color": "#56ABFB"
}
部分更新一个工单状态
PATCH
https://rest_api_root/v1/ship/ticket_states/{ticket_state_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
ticket_state_id	String	
工单状态id。

请求参数
字段	类型	描述
name可选	String	
工单状态的名称，在一个企业中这个名称是唯一的。

type可选	String	
工单状态的类型。

允许值: pending, in_progress, completed, closed

请求示例：
{
    "name": "已完成",
    "type": "completed"
}
响应示例：
{
    "id": "6422711c3f12e6c1e46d40f2",
    "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
    "name": "已完成",
    "type": "completed",
    "color": "#56AB25"
}
获取全部工单状态列表
GET
https://rest_api_root/v1/ship/ticket_states
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40f2",
            "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
            "name": "处理中",
            "type": "pending",
            "color": "#F6C659"
        }
    ]
}
获取工单状态方案列表
GET
https://rest_api_root/v1/ship/ticket_state_plans
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63feb3da9cc1ead1d2be93f4",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4",
            "product": null
        },
        {
            "id": "63feb3da9cc1ead1d2be93f5",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f5",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向状态方案中添加一个工单状态
POST
https://rest_api_root/v1/ship/ticket_state_plans/{state_plan_id}/ticket_states
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工单状态方案的id。

请求参数
字段	类型	描述
state_id	String	
工单状态的id。

请求示例：
{
    "state_id": "63bb744214bd13c9def24ca2"
}
响应示例：
{
    "id": "63bb744214bd13c9def24ca2",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63bb744214bd13c9def24ca2",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
获取状态方案中的工单状态列表
GET
https://rest_api_root/v1/ship/ticket_state_plans/{state_plan_id}/ticket_states
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工单状态方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca2",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63bb744214bd13c9def24ca2",
            "state_plan": {
                "id": "63feb3da9cc1ead1d2be93f4",
                "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
            },
            "state": {
                "id": "63bb744214bd13c9def24ca2",
                "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
                "name": "待处理",
                "type": "pending"
            }
        }
    ]
}
在状态方案中移除一个工单状态
移除状态后，每种类型的状态至少存在一种，否则将无法移除。

DELETE
https://rest_api_root/v1/ship/ticket_state_plans/{state_plan_id}/ticket_states/{state_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工单状态方案的id。

state_id	String	
工单状态的id。

响应示例：
{
    "id": "63bb744214bd13c9def24ca2",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63bb744214bd13c9def24ca2",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
向状态方案中添加一个工单状态流转
POST
https://rest_api_root/v1/ship/ticket_state_plans/{state_plan_id}/ticket_state_flows
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工单状态方案的id。

请求参数
字段	类型	描述
from_state_id	String	
起始工单状态的id。

to_state_id	String	
可更改为的工单状态的id。

请求示例：
{
    "from_state_id": "63bb744214bd13c9def24ca5",
    "to_state_id": "63bb744214bd13c9def24ca2"
}
响应示例：
{
    "id": "63feb3da9cc1ead1d2be93fd",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_state_flows/63feb3da9cc1ead1d2be93fd",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "form_state": {
        "id": "63bb744214bd13c9def24ca5",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca5",
        "name": "已计划",
        "type": "in_progress"
    },
    "to_state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
获取状态方案中的工单状态流转列表
GET
https://rest_api_root/v1/ship/ticket_state_plans/{state_plan_id}/ticket_state_flows
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工单状态方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63feb3da9cc1ead1d2be93f5",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63feb3da9cc1ead1d2be93f5",
            "state_plan": {
                "id": "63feb3da9cc1ead1d2be93f4",
                "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
            },
            "form_state": {
                "id": "63bb744214bd13c9def24ca2",
                "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
                "name": "待处理",
                "type": "pending"
            },
            "to_state": {
                "id": "63bb744214bd13c9def24ca4",
                "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca4",
                "name": "处理中",
                "type": "in_progress"
            }
        }
    ]
}
在状态方案中移除一个工单状态流转
DELETE
https://rest_api_root/v1/ship/ticket_state_plans/{state_plan_id}/ticket_state_flows/{state_flow_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工单状态方案的id。

state_flow_id	String	
工单状态流转的id。

响应示例：
{
    "id": "63feb3da9cc1ead1d2be93fd",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_state_flows/63feb3da9cc1ead1d2be93fd",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "form_state": {
        "id": "63bb744214bd13c9def24ca5",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca5",
        "name": "已计划",
        "type": "in_progress"
    },
    "to_state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
创建一个工单属性
POST
https://rest_api_root/v1/ship/ticket_properties
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
工单属性的名称。在一个企业中这个名称是唯一的。

type	String	
工单属性的类型。

允许值: text, textarea, select, multi_select, cascade_select, cascade_multi_select, member, members, date, number, progress, rate, link

options可选	Object[]	
选择项列表。当工单属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例（select）：
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
请求示例（cascade_select）：
{
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ]
}
响应示例（select）：
{
    "id": "severity",
    "url": "https://rest_api_root/v1/ship/ticket_properties/severity",
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
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
响应示例（cascade_select）：
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/ticket_properties/jiliandanxuan",
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
部分更新一个工单属性
PATCH
https://rest_api_root/v1/ship/ticket_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_id	String	
工单属性的id。

请求参数
字段	类型	描述
name可选	String	
工单属性的名称。在一个企业中这个名称是唯一的。

options可选	Object[]	
选择项列表。options是整体更新的。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例（select）：
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
请求示例（cascade_select）：
{
    "name": "级联单选-update",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ]
}
响应示例（select）：
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/ship/ticket_properties/severity",
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
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
响应示例（cascade_select）：
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/ticket_properties/jiliandanxuan",
    "name": "级联单选-update",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
获取全部工单属性列表
GET
https://rest_api_root/v1/ship/ticket_properties
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
            "name": "解决方案",
            "type": "select",
            "options": [
                {
                    "_id": "6422711c3f12e6c1e46d40e9",
                    "text": "进入需求池"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/ticket_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
获取工单属性方案列表
GET
https://rest_api_root/v1/ship/ticket_property_plans
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5f8a21f18ef715265de90c21",
            "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21",
            "product": null
        },
        {
            "id": "5f8a21f18ef715265de90c22",
            "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c22",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向工单属性方案中添加一个工单属性
POST
https://rest_api_root/v1/ship/ticket_property_plans/{property_plan_id}/ticket_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
工单属性方案的id。

请求参数
字段	类型	描述
property_id	String	
工单属性的id。

请求示例：
{
    "property_id": "solution"
}
响应示例：
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21/ticket_properties/solution",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
获取工单属性方案中的工单属性列表
GET
https://rest_api_root/v1/ship/ticket_property_plans/{property_plan_id}/ticket_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
工单属性方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21/ticket_properties/solution",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21"
            },
            "property": {
                "id": "solution",
                "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
                "name": "解决方案",
                "type": "select"
            }
        }
    ]
}
在工单属性方案中移除一个工单属性
DELETE
https://rest_api_root/v1/ship/ticket_property_plans/{property_plan_id}/ticket_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
工单属性方案的id。

property_id	String	
工单属性的id。

响应示例：
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21/ticket_properties/solution",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
需求配置
用于查看和管理需求相关的配置信息。

获取全部需求状态列表
GET
https://rest_api_root/v1/ship/idea_states
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63e1bf51898a0be5a2d21b2a",
            "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
            "name": "待评审",
            "type": "pending",
            "color": "#56ABFB"
        }
    ]
}
创建一个需求属性
POST
https://rest_api_root/v1/ship/idea_properties
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
需求属性的名称。在一个企业中这个名称是唯一的。

type	String	
需求属性的类型。

允许值: text, textarea, select, multi_select, cascade_select, cascade_multi_select, member, members, date, number, progress, rate, link

options可选	Object[]	
选择项列表。当需求属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例（select）：
{
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重"
        },
        {
            "text": "一般"
        }
    ]
}
请求示例（cascade_select）：
{
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ]
}
响应示例（select）：
{
    "id": "severity",
    "url": "https://rest_api_root/v1/ship/idea_properties/severity",
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
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
响应示例（cascade_select）：
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/idea_properties/jiliandanxuan",
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
部分更新一个需求属性
PATCH
https://rest_api_root/v1/ship/idea_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_id	String	
需求属性的id。

请求参数
字段	类型	描述
name可选	String	
需求属性的名称。在一个企业中这个名称是唯一的。

options可选	Object[]	
选择项列表。options是整体更新的。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例（select）：
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
请求示例（cascade_select）：
{
    "name": "级联单选-update",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子-update",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ]
}
响应示例（select）：
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/ship/idea_properties/severity",
    "name": "严重程度-update",
    "type": "select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
响应示例（cascade_select）：
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/idea_properties/jiliandanxuan",
    "name": "级联单选-update",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子-update",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
获取全部需求属性列表
GET
https://rest_api_root/v1/ship/idea_properties
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "backlog_type",
            "url": "https://rest_api_root/v1/ship/idea_properties/backlog_type",
            "name": "需求类型",
            "type": "select",
            "options": [
                {
                    "_id": "5cb7e763fda1ce4ca0010002",
                    "text": "功能需求"
                },
                {
                    "_id": "5cb7e763fda1ce4ca0010004",
                    "text": "体验优化"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/idea_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
获取需求属性方案列表
GET
https://rest_api_root/v1/ship/idea_property_plans
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5d7a21f19ef715269ae90c50",
            "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50",
            "product": null
        },
        {
            "id": "5f8a21f18ef715265de90c22",
            "url": "https://rest_api_root/v1/ship/idea_property_plans/5f8a21f18ef715265de90c22",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向需求属性方案中添加一个需求属性
POST
https://rest_api_root/v1/ship/idea_property_plans/{property_plan_id}/idea_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
需求属性方案的id。

请求参数
字段	类型	描述
property_id	String	
需求属性的id。

请求示例：
{
    "property_id": "solution"
}
响应示例：
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50/idea_properties/solution",
    "property_plan": {
        "id": "5d7a21f19ef715269ae90c50",
        "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/idea_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
获取需求属性方案中的需求属性列表
GET
https://rest_api_root/v1/ship/idea_property_plans/{property_plan_id}/idea_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
需求属性方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50/idea_properties/solution",
            "property_plan": {
                "id": "5d7a21f19ef715269ae90c50",
                "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50"
            },
            "property": {
                "id": "solution",
                "url": "https://rest_api_root/v1/ship/idea_properties/solution",
                "name": "解决方案",
                "type": "select"
            }
        }
    ]
}
在需求属性方案中移除一个需求属性
DELETE
https://rest_api_root/v1/ship/idea_property_plans/{property_plan_id}/idea_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
需求属性方案的id。

property_id	String	
需求属性的id。

响应示例：
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50/idea_properties/solution",
    "property_plan": {
        "id": "5d7a21f19ef715269ae90c50",
        "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/idea_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}