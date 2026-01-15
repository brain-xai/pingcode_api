项目管理
项目
用于查看和管理项目相关信息。
资源地址：{GET} https://rest_api_root/v1/project/projects/{project_id}

资源属性
字段	类型	描述
id	String	
项目的id。

url	String	
项目的地址。

identifier	String	
项目的标识。

name	String	
项目的名称。

type	String	
项目的类型。

允许值: kanban, scrum, waterfall, hybrid

assignee	Object	
项目的负责人。

scope_type	String	
项目的所属类型。

允许值: organization, user_group

scope_id	String	
项目的所属id。

visibility	String	
项目的可见性。

允许值: private, public

state	Object	
项目的状态。

start_at	Number	
项目的开始时间。

end_at	Number	
项目的结束时间。

color	String	
项目的主题色。

description	String	
项目的描述。

properties	Object	
项目的自定义属性。

properties.prop_a	Object	
项目的自定义属性prop_a。

properties.prop_b	Object	
项目的自定义属性prop_b。

members	Object[]	
项目的成员列表。

created_at	Number	
项目的创建时间。

created_by	Object	
项目的创建人。

updated_at	Number	
项目的更新时间。

updated_by	Object	
项目的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "color": "#F693E7",
    "description": "这是一个scrum类型的项目",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
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
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/63c8fb32729dee3334d96af7",
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
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "is_archived": 0,
    "is_deleted": 0
}
创建一个项目
POST
https://rest_api_root/v1/project/projects
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
scope_type可选	String	
项目的所属类型。默认值organization。允许值分别表示企业可见和团队可见。

默认值: organization

允许值: organization, user_group

scope_id可选	String	
项目的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。

name	String	
项目的名称。最大长度为255。

visibility可选	String	
项目的可见范围。允许值分别表示组织全部成员可见和仅项目成员可见。

默认值: private

允许值: public, private

type	String	
项目的类型。

允许值: kanban, scrum, waterfall, hybrid

identifier	String	
项目的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
项目的描述。

members可选	Object[]	
项目成员的列表。当项目的scope_type变为user_group时，项目成员必须是scope_id指定的团队内的成员；当scope_type为其他值时，项目成员可以是任意成员或团队。

members.id	String	
企业成员或团队的id。

members.type	String	
项目成员的类型。

允许值: user, user_group

start_at可选	Number	
项目开始的时间。

end_at可选	Number	
项目结束的时间。

assignee_id可选	String	
项目负责人的id。

请求示例：
{
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "Scrum项目",
    "visibility": "private",
    "type": "scrum",
    "identifier": "SCR",
    "description": "这是一个scrum类型的项目",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "type": "user"
        }
    ],
    "start_at": 1680278400,
    "end_at": 1682870399,
    "assignee_id": "8168dd057b104c81bceb2e48bdf283d0"
}
响应示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "color": "#F693E7",
    "description": "这是一个scrum类型的项目",
    "properties": {},
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
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
复制一个项目
POST
https://rest_api_root/v1/project/projects/{project_id}/clone
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
scope_type可选	String	
项目的所属类型。默认使用原项目的所属。允许值分别表示企业可见和团队可见。

允许值: organization, user_group

scope_id可选	String	
项目的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。

name可选	String	
项目的名称。最大长度为255。默认使用原项目的名称。

visibility可选	String	
项目的可见范围。默认使用原项目的可见范围。允许值分别表示组织全部成员可见和仅项目成员可见。

允许值: public, private

identifier	String	
项目的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
项目的描述。默认使用原项目的描述。

members可选	Object[]	
项目成员的列表。当项目的scope_type变为user_group时，项目成员必须是scope_id指定的团队内的成员；当scope_type为其他值时，项目成员可以是任意成员或团队，默认使用原项目的成员列表。

members.id	String	
企业成员或团队的id。

members.type	String	
项目成员的类型。

允许值: user, user_group

请求示例：
{
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "复制的Scrum项目",
    "visibility": "public",
    "identifier": "SCRC",
    "description": "这是一个复制的Scrum类型的项目",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "type": "user"
        }
    ]
}
响应示例：
{
    "id": "5ab623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5ab623f6a70571487ea47001",
    "identifier": "SCRC",
    "name": "复制的Scrum项目",
    "type": "scrum",
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "public",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "color": "#F693E7",
    "description": "这是一个复制的Scrum类型的项目",
    "properties": {},
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
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
部分更新一个项目
PATCH
https://rest_api_root/v1/project/projects/{project_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name可选	String	
项目的名称。最大长度为255。

identifier可选	String	
项目的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。

description可选	String	
项目的描述。

start_at可选	Number	
项目开始的时间。

end_at可选	Number	
项目结束的时间。

assignee_id可选	String	
项目负责人的id。

state_id可选	String	
项目状态的id。项目状态可以通过获取项目状态列表获取。

properties可选	Object	
项目自定义属性的键值对集合。需要注意自定义属性需要包含在项目属性方案中。例如属性方案中包含prop_a和prop_b。

properties.prop_a可选	Object	
项目自定义属性prop_a。

properties.prop_b可选	Object	
项目自定义属性prop_b。

请求示例：
{
    "name": "Scrum项目",
    "identifier": "SCR",
    "description": "这是一个scrum类型的项目",
    "start_at": 1680278400,
    "end_at": 1682870399,
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "state_id": "66cbf3b4b78a55fcd1a76296",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "color": "#F693E7",
    "description": "这是一个scrum类型的项目",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
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
获取项目列表
GET
https://rest_api_root/v1/project/projects
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
identifier可选	String	
项目的标识。

type可选	String	
项目的类型。

允许值: scrum, kanban, waterfall, hybrid

include_deleted可选	Boolean	
是否查询已删除的项目。该值默认为false。

include_archived可选	Boolean	
是否查询已归档的项目。该值默认为false。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47000",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
            "identifier": "SCR",
            "name": "Scrum项目",
            "type": "scrum",
            "assignee": {
                "id": "8168dd057b104c81bceb2e48bdf283d0",
                "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
                "name": "john",
                "display_name": "john",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "scope_type": "user_group",
            "scope_id": "63c8fb32729dee3334d96af7",
            "visibility": "private",
            "state": {
                "id": "66cbf3b4b78a55fcd1a76296",
                "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
                "name": "正常",
                "type": "in_progress"
            },
            "start_at": 1680278400,
            "end_at": 1682870399,
            "color": "#F693E7",
            "description": "这是一个scrum类型的项目",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
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
获取项目状态列表
GET
https://rest_api_root/v1/project/project/states?project_id={project_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 5,
    "values": [
        {
            "id": "66cbf5401e7cc374c85acb1b",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1b",
            "name": "未开始",
            "type": "pending"
        },
        {
            "id": "66cbf5401e7cc374c85acb1c",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1c",
            "name": "正常",
            "type": "in_progress"
        },
        {
            "id": "66cbf5401e7cc374c85acb1d",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1d",
            "name": "预警",
            "type": "in_progress"
        },
        {
            "id": "66cbf5401e7cc374c85acb1e",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1e",
            "name": "延期",
            "type": "in_progress"
        },
        {
            "id": "66cbf5401e7cc374c85acb1f",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1f",
            "name": "结束",
            "type": "completed"
        }
    ]
}
向项目中添加一个成员
POST
https://rest_api_root/v1/project/projects/{project_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
member	Object	
项目的成员。

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
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
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
部分更新一个项目内的成员
PATCH
https://rest_api_root/v1/project/projects/{project_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

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
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
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
获取项目中的成员列表
GET
https://rest_api_root/v1/project/projects/{project_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
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
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/63c8fb32729dee3334d96af7",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
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
在项目中移除一个成员
DELETE
https://rest_api_root/v1/project/projects/{project_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

member_id	String	
企业成员或团队的id。

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
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
向项目中添加一个项目属性
POST
https://rest_api_root/v1/project/projects/{project_id}/project_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
property_id	String	
项目属性的id。

请求示例：
{
    "property_id": "risk"
}
响应示例：
{
    "id": "risk",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/risk",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "property": {
        "id": "risk",
        "url": "http://rest_api_root/v1/project/project_properties/risk",
        "name": "项目风险",
        "type": "select",
        "options": [
            {
                "_id": "5efb1859110533727a82c603",
                "text": "高"
            },
            {
                "_id": "5efb1859110533727a82c604",
                "text": "中"
            },
            {
                "_id": "5efb1859110533727a82c605",
                "text": "低"
            }
        ]
    }
}
获取项目中的项目属性列表
GET
https://rest_api_root/v1/project/projects/{project_id}/project_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "risk",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/risk",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "property": {
                "id": "risk",
                "url": "http://rest_api_root/v1/project/project_properties/risk",
                "name": "项目风险",
                "type": "select",
                "options": [
                    {
                        "_id": "5efb1859110533727a82c603",
                        "text": "高"
                    },
                    {
                        "_id": "5efb1859110533727a82c604",
                        "text": "中"
                    },
                    {
                        "_id": "5efb1859110533727a82c605",
                        "text": "低"
                    }
                ]
            }
        },
        {
            "id": "name",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/name",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "property": {
                "id": "name",
                "url": "https://rest_api_root/v1/project/project_properties/name",
                "name": "名称",
                "type": "text",
                "options": null,
                "is_removable": false,
                "is_name_editable": false,
                "is_options_editable": false
            }
        }
    ]
}
在项目中移除一个项目属性
DELETE
https://rest_api_root/v1/project/projects/{project_id}/project_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

property_id	String	
项目属性的id。

响应示例：
{
    "id": "risk",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/risk",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "property": {
        "id": "risk",
        "url": "http://rest_api_root/v1/project/project_properties/risk",
        "name": "项目风险",
        "type": "select",
        "options": [
            {
                "_id": "5efb1859110533727a82c603",
                "text": "高"
            },
            {
                "_id": "5efb1859110533727a82c604",
                "text": "中"
            },
            {
                "_id": "5efb1859110533727a82c605",
                "text": "低"
            }
        ]
    }
}
获取一个项目进度
GET
https://rest_api_root/v1/project/projects/{project_id}/progress
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "work_item": {
        "total": 4,
        "pending_count": 1,
        "in_progress_count": 2,
        "completed_count": 1
    }
}
Scrum
用于查看和管理迭代相关信息。

创建一个迭代
POST
https://rest_api_root/v1/project/projects/{project_id}/sprints
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
迭代的名称。

start_at	Number	
迭代开始的时间。

end_at	Number	
迭代结束的时间。

assignee_id	String	
迭代负责人的id。

description可选	String	
迭代的描述。

status可选	String	
迭代的状态。

允许值: pending, in_progress, completed

category_ids可选	String[]	
迭代类别的id数组。

请求示例：
{
    "name": "Sprint 2",
    "start_at": 1589791860,
    "end_at": 1589791860,
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "description": "This is sprint 2",
    "status": "pending",
    "category_ids": [
        "676a460a0fd987b7ea320887"
    ]
}
响应示例：
{
    "id": "5ecf7b74eaab845a2aa53132",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53132",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Sprint 2",
    "status": "pending",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1589791860,
    "end_at": 1589791860,
    "description": "This is sprint 2",
    "started_at": 1589791860,
    "completed_at": 1589791960,
    "total_story_points": 0,
    "started_story_points": 0,
    "completed_story_points": 0,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320887",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
            "name": "Category 1"
        }
    ],
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
    }
}
批量创建迭代
POST
https://rest_api_root/v1/project/sprints/bulk
权限: 企业令牌

请求参数
字段	类型	描述
sprints	Object[]	
需要批量创建的迭代。该参数是一个对象数组（数组内对象不得超过100个）。

sprints.project_id	String	
迭代所属项目的id。

sprints.name	String	
迭代的名称。

sprints.start_at	Number	
迭代开始的时间。

sprints.end_at	Number	
迭代结束的时间。

sprints.assignee_id	String	
迭代负责人的id。

sprints.description可选	String	
迭代的描述。

sprints.status可选	String	
迭代的状态。

允许值: pending, in_progress, completed

sprints.category_ids可选	String[]	
迭代类别的id列表。

请求示例：
{
    "sprints": [
        {
            "project_id": "5eb623f6a70571487ea47000",
            "name": "Sprint 3",
            "start_at": 1589791860,
            "end_at": 1589791860,
            "assignee_id": "a0417f68e846aae315c85d24643678a9",
            "description": "This is sprint 3",
            "status": "pending",
            "category_ids": [
                "5e6b35de50ef8153c2062f70"
            ]
        }
    ]
}
响应示例：
[
    {
        "state": "success",
        "sprint": {
            "id": "5ecf7b74eaab845a2aa53134",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53134",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Sprint 3",
            "status": "pending",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1589791860,
            "end_at": 1589791860,
            "description": "This is sprint 3",
            "started_at": 1589791860,
            "completed_at": 1589791960,
            "total_story_points": 0,
            "started_story_points": 0,
            "completed_story_points": 0,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320887",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
                    "name": "Category 1"
                }
            ],
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
            }
        }
    }
]
部分更新一个迭代
PATCH
https://rest_api_root/v1/project/projects/{project_id}/sprints/{sprint_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

sprint_id	String	
迭代的id。

请求参数
字段	类型	描述
name可选	String	
迭代的名称。

start_at可选	Number	
迭代开始的时间。

end_at可选	Number	
迭代结束的时间。

assignee_id可选	String	
迭代负责人的id。

description可选	String	
迭代的描述。

status可选	String	
迭代的状态。

允许值: pending, in_progress, completed

category_ids可选	String[]	
迭代类别的id列表。

请求示例：
{
    "name": "Sprint 2",
    "start_at": 1589791860,
    "end_at": 1589791860,
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "description": "This is sprint 2",
    "status": "in_progress",
    "category_ids": [
        "5e6b35de50ef8153c2062f70"
    ]
}
响应示例：
{
    "id": "5ecf7b74eaab845a2aa53132",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53132",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Sprint 2",
    "status": "in_progress",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1589791860,
    "end_at": 1589791860,
    "description": "This is sprint 2",
    "started_at": 1589791860,
    "completed_at": 1589791960,
    "total_story_points": 0,
    "started_story_points": 0,
    "completed_story_points": 0,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320887",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
            "name": "Category 1"
        }
    ],
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
    }
}
获取迭代列表
GET
https://rest_api_root/v1/project/projects/{project_id}/sprints
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

查询参数
字段	类型	描述
name可选	String	
迭代的名称。

status可选	String	
迭代的状态。

允许值: pending, in_progress, completed

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
            "id": "5ecf7b74eaab845a2aa53138",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Sprint 1",
            "status": "completed",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1589791860,
            "end_at": 1589791860,
            "description": "This is sprint 1",
            "started_at": 1589791860,
            "completed_at": 1589791960,
            "total_story_points": 0,
            "started_story_points": 0,
            "completed_story_points": 0,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320887",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
                    "name": "Category 1"
                }
            ],
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
            }
        }
    ]
}
创建一个迭代分组
POST
https://rest_api_root/v1/project/projects/{project_id}/sprint_sections
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
迭代分组的名称。

请求示例：
{
    "name": "Section 1"
}
响应示例：
{
    "id": "634f869a0fd987b7ea320833",
    "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320833",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Section 1"
}
部分更新一个迭代分组
PATCH
https://rest_api_root/v1/project/projects/{project_id}/sprint_sections/{section_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

section_id	String	
迭代分组的id。

请求参数
字段	类型	描述
name	String	
迭代分组的名称。

请求示例：
{
    "name": "Section 1"
}
响应示例：
{
    "id": "634f869a0fd987b7ea320833",
    "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320833",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Section 1"
}
获取迭代分组列表
GET
https://rest_api_root/v1/project/projects/{project_id}/sprint_sections
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "634f869a0fd987b7ea320833",
            "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320833",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Section 1"
        },
        {
            "id": "634f869a0fd987b7ea320834",
            "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320834",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Section 2"
        }
    ]
}
删除一个迭代分组
DELETE
https://rest_api_root/v1/project/projects/{project_id}/sprint_sections/{section_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

section_id	String	
迭代分组的id。

响应示例：
{
    "id": "634f869a0fd987b7ea320834",
    "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320834",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Section 2"
}
创建一个迭代类别
POST
https://rest_api_root/v1/project/projects/{project_id}/sprint_categories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
迭代类别的名称。

section_id可选	String	
迭代类别所属迭代分组id。

请求示例：
{
    "name": "Category 1",
    "section_id": "634f869a0fd987b7ea320833"
}
响应示例：
{
    "id": "676a460a0fd987b7ea320887",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Category 1",
    "section": {
        "id": "634f869a0fd987b7ea320833",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
        "name": "Section 1"
    }
}
部分更新一个迭代类别
PATCH
https://rest_api_root/v1/project/projects/{project_id}/sprint_categories/{sprint_category_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

sprint_category_id	String	
迭代类别的id。

请求参数
字段	类型	描述
name可选	String	
迭代类别的名称。

section_id可选	String	
迭代类别所属迭代分组id。

请求示例：
{
    "name": "Category 2",
    "section_id": "634f869a0fd987b7ea320833"
}
响应示例：
{
    "id": "676a460a0fd987b7ea320888",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320888",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Category 2",
    "section": {
        "id": "634f869a0fd987b7ea320833",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
        "name": "Section 1"
    }
}
获取迭代类别列表
GET
https://rest_api_root/v1/project/projects/{project_id}/sprint_categories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "676a460a0fd987b7ea320887",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Category 1",
            "section": {
                "id": "634f869a0fd987b7ea320833",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
                "name": "Section 1"
            }
        },
        {
            "id": "676a460a0fd987b7ea320888",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320888",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Category 2",
            "section": {
                "id": "634f869a0fd987b7ea320833",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
                "name": "Section 1"
            }
        }
    ]
}
删除一个迭代类别
DELETE
https://rest_api_root/v1/project/projects/{project_id}/sprint_categories/{sprint_category_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

sprint_category_id	String	
迭代类别的id。

响应示例：
{
    "id": "676a460a0fd987b7ea320888",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320888",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Category 2",
    "section": {
        "id": "634f869a0fd987b7ea320833",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
        "name": "Section 1"
    }
}
Kanban
用于查看和管理看板相关信息。

创建一个看板
POST
https://rest_api_root/v1/project/projects/{project_id}/boards
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
看板的名字。在同一个项目中该名字是唯一的。

work_item_types可选	String[]	
看板支持的工作项类型列表。自定义工作项类型只支持hybrid类型项目里的看板。

允许值: epic, feature, story, task, bug, issue, 自定义工作项类型id

请求示例：
{
    "name": "一个看板",
    "work_item_types": [
        "epic",
        "feature",
        "story",
        "6385c650fef18f2d7222d15d"
    ]
}
响应示例：
{
    "id": "5eb623f6a70571487ea47222",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "一个看板",
    "work_item_types": [
        "epic",
        "feature",
        "story",
        "6385c650fef18f2d7222d15d"
    ]
}
部分更新一个看板
PATCH
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

请求参数
字段	类型	描述
name可选	String	
看板的名字。在同一个项目中该名字是唯一的。

work_item_types可选	String[]	
看板支持的工作项类型列表。

允许值: epic, feature, story, task, bug, issue, 自定义工作项类型id

请求示例：
{
    "name": "一个看板",
    "work_item_types": [
        "epic",
        "feature",
        "6385c650fef18f2d7222d15d"
    ]
}
响应示例：
{
    "id": "5eb623f6a70571487ea47222",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "一个看板",
    "work_item_types": [
        "epic",
        "feature",
        "6385c650fef18f2d7222d15d"
    ]
}
获取看板列表
GET
https://rest_api_root/v1/project/projects/{project_id}/boards
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47222",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "kanban",
            "work_item_types": [
                "epic",
                "feature",
                "story"
            ]
        }
    ]
}
删除一个看板
DELETE
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

响应示例：
{
    "id": "5eb623f6a70571487ea47222",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "一个看板",
    "work_item_types": [
        "epic",
        "feature"
    ]
}
创建一个看板栏
POST
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/entries
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

请求参数
字段	类型	描述
name	String	
看板栏的名称。在同一看板下该名称是唯一的。

wip_limit可选	Number	
在制品数量。

is_split可选	Boolean	
是否将看板栏拆分为进行中和已完成。

默认值: false

definition_of_done可选	String	
完成的定义。

请求示例：
{
    "name": "一个看板栏",
    "wip_limit": 1,
    "is_split": true,
    "definition_of_done": "Unit test passed"
}
响应示例：
{
    "id": "5ab623f6a70571487ea45634",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222/entries/5ab623f6a70571487ea45634",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47222",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
        "name": "kanban",
        "work_item_types": [
            "epic",
            "feature",
            "story"
        ]
    },
    "name": "一个看板栏",
    "is_system": false,
    "is_split": true,
    "wip_limit": 1,
    "definition_of_done": "Unit test passed"
}
部分更新一个看板栏
PATCH
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/entries/{entry_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

entry_id	String	
看板栏的id。

请求参数
字段	类型	描述
name可选	String	
看板栏的名称。在同一看板下该名称是唯一的。

wip_limit可选	Number	
在制品数量。

is_split可选	Boolean	
是否将看板栏拆分为进行中和已完成

默认值: false

definition_of_done可选	String	
完成的定义。

请求示例：
{
    "name": "需求池",
    "wip_limit": 1,
    "is_split": true,
    "definition_of_done": "Unit test passed"
}
响应示例：
{
    "id": "5ab623f6a70571487ea45634",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/entries/5ab623f6a70571487ea45634",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": [
            "epic",
            "feature",
            "story"
        ]
    },
    "name": "需求池",
    "is_system": false,
    "is_split": true,
    "wip_limit": 1,
    "definition_of_done": "Unit test passed"
}
获取看板栏列表
GET
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/entries
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ab623f6a70571487ea45634",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/entries/5ab623f6a70571487ea45634",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "board": {
                "id": "5eb623f6a70571487ea47223",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
                "name": "默认看板",
                "work_item_types": [
                    "epic",
                    "feature",
                    "story"
                ]
            },
            "name": "需求池",
            "is_system": false,
            "is_split": true,
            "wip_limit": 1,
            "definition_of_done": "Unit test passed"
        }
    ]
}
删除一个看板栏
DELETE
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/entries/{entry_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

entry_id	String	
看板栏的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ab623f6a70571487ea45634",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/entries/5ab623f6a70571487ea45634",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "board": {
                "id": "5eb623f6a70571487ea47223",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
                "name": "默认看板",
                "work_item_types": [
                    "epic",
                    "feature",
                    "story"
                ]
            },
            "name": "需求池",
            "is_system": false,
            "is_split": true,
            "wip_limit": 1,
            "definition_of_done": "Unit test passed"
        }
    ]
}
创建一个泳道
POST
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/swimlanes
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

请求参数
字段	类型	描述
name	String	
泳道的名称。在同一看板下该名称是唯一的。

请求示例：
{
    "name": "一个泳道"
}
响应示例：
{
    "id": "5bb623f6a70571487ea44357",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": [
            "epic",
            "feature",
            "story"
        ]
    },
    "name": "一个泳道",
    "is_system": false
}
部分更新一个泳道
PATCH
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/swimlanes/{swimlane_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

swimlane_id	String	
泳道的id。

请求参数
字段	类型	描述
name可选	String	
泳道的名称。在同一看板下该名称是唯一的。

请求示例：
{
    "name": "一个泳道"
}
响应示例：
{
    "id": "5bb623f6a70571487ea44357",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": [
            "epic",
            "feature",
            "story"
        ]
    },
    "name": "一个泳道",
    "is_system": false
}
获取泳道列表
GET
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/swimlanes
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5bb623f6a70571487ea44357",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "board": {
                "id": "5eb623f6a70571487ea47223",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
                "name": "默认看板",
                "work_item_types": [
                    "epic",
                    "feature",
                    "story"
                ]
            },
            "name": "一个泳道",
            "is_system": false
        }
    ]
}
删除一个泳道
DELETE
https://rest_api_root/v1/project/projects/{project_id}/boards/{board_id}/swimlanes/{swimlane_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

board_id	String	
看板的id。

swimlane_id	String	
泳道的id.

响应示例：
{
    "id": "5bb623f6a70571487ea44357",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": [
            "epic",
            "feature",
            "story"
        ]
    },
    "name": "一个泳道",
    "is_system": false
}
工作项
用于查看和管理工作项相关信息。
资源地址：{GET} https://rest_api_root/v1/project/work_items/{work_item_id}

资源属性
字段	类型	描述
id	String	
工作项的id。

url	String	
工作项的地址。

project	Object	
工作项所属的项目。

identifier	String	
工作项的标识。

title	String	
工作项的标题。

type	String	
工作项的类型。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。

start_at	Number	
工作项的开始时间。

end_at	Number	
工作项的结束时间。

parent_id	String	
工作项的父工作项id。

short_id	String	
工作项的短id。

html_url	String	
工作项的html地址。

parent	Object	
工作项的父工作项。

assignee	Object	
工作项的负责人。

state	Object	
工作项的状态。

priority	Object	
工作项的优先级。

version	Object	
工作项属的发布。

sprint	Object	
工作项所属的迭代。

board	Object	
工作项所属的看板。

entry	Object	
工作项所属的看板栏。

swimlane	Object	
工作项的所属的泳道。

phase	Object	
工作项的所属计划。

description	String	
工作项的描述。

completed_at	Number	
工作项的完成时间。

story_points	Number	
工作项的故事点。

estimated_workload	Number	
工作项的预估工时。

remaining_workload	Number	
工作项的剩余工时。

properties	Object	
工作项的自定义属性。自定义属性包括用户自定义的属性和一些系统内置的属性。用户自定义的属性例如prop_a和prop_b。系统内置的属性例如bug_type、solution和risk等。

tags	Object[]	
工作项的标签列表。

participants	Object[]	
工作项的关注人列表。

public_image_token	String	
工作项描述和自定义多行文本类型属性里获取图片资源token。获取一个工作项和获取工作项列表参数include_public_image_token值有效时返回。

created_at	Number	
工作项的创建时间。

created_by	Object	
工作项的创建人。

updated_at	Number	
工作项的更新时间。

updated_by	Object	
工作项的更新人。

is_archived	Number	
是否已归档。

允许值: 0, 1

is_deleted	Number	
是否已删除。

允许值: 0, 1

全量数据示例：
{
    "id": "6efca131b06329c524cdd2fb",
    "url": "https://rest_api_root/v1/project/work_items/6efca131b06329c524cdd2fb",
    "project": {
        "id": "5eb623f6a70571487ea47004",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004",
        "identifier": "HBR",
        "name": "Hybrid项目",
        "type": "hybrid",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "HBR-1",
    "title": "这是一个用户故事",
    "type": "story",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": "5edfa3b67463c1ee626c0980",
    "short_id": "eqWqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/eqWqLmTO",
    "parent": {
        "id": "5edfa3b67463c1ee626c0980",
        "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0980",
        "identifier": "HBR-2",
        "title": "这是一个特性",
        "type": "feature",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": null,
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value",
            "risk": null,
            "business_value": null,
            "effort": null,
            "backlog_type": null,
            "backlog_from": null
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/versions/5eb623487ea47001f6a70582",
        "name": "1.0.1",
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
    "sprint": {
        "id": "5ecf7b74eaab845a2aa53139",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/sprints/5ecf7b74eaab845a2aa53139",
        "name": "Sprint 1",
        "start_at": 1589791860,
        "end_at": 1589791860,
        "status": "completed"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board": {
        "id": "5eb623f6a70571487ea47333",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333",
        "name": "kanban",
        "work_item_types": [
            "epic",
            "feature",
            "issue",
            "story"
        ]
    },
    "entry": {
        "id": "5ee1c4fac5b3c51206f0a862",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/entries/5ee1c4fac5b3c51206f0a862",
        "name": "需求池"
    },
    "swimlane": {
        "id": "5ee1c4fac5b3c51206f0a867",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/swimlanes/5ee1c4fac5b3c51206f0a867",
        "name": "默认泳道"
    },
    "phase": {
        "id": "63761fee31caaf77189816b5",
        "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/phases/63761fee31caaf77189816b5",
        "title": "这是一个阶段",
        "identifier": "WTF-1"
    },
    "description": "<p>这是一个用户故事的描述<p><img src=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" originUrl=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" alt=\"图片.png\" size=\"35460\" style=\"text-align: center;\" />",
    "completed_at": 1583290347,
    "story_points": 1,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "risk": null,
        "backlog_type": null,
        "backlog_from": null
    },
    "tags": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
            "name": "标签-1"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
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
    "public_image_token": "-fkvANQ2dcVECK6Xg45L3kG8VCbOTK8NrNckGkxljRY",
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
    "id": "5edca5d2fa112b06305c24ca",
    "url": "https://rest_api_root/v1/project/work_items/5edca5d2fa112b06305c24ca",
    "identifier": "KB-1",
    "title": "这是一个事务",
    "type": "issue",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": null,
    "short_id": "c9WqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "entry_status": null,
        "entry_position": null,
        "operation_time": 1691571221
    }
}
创建一个工作项
POST
https://rest_api_root/v1/project/work_items
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
project_id	String	
项目的id。

type_id	String	
工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。

title	String	
工作项的标题。

description可选	String	
工作项的描述。

start_at可选	Number	
工作项的开始时间。当工作项类型为里程碑时，该参数无效。

end_at可选	Number	
工作项的截止时间。

priority_id可选	String	
工作项优先级的id。

state_id可选	String	
工作项状态的id。工作项状态的id在设置时必须同时满足工作项类型的工作项状态方案和工作项状态流转的状态值才能成功完成设置。工作项状态方案可以通过获取工作项状态方案列表获取。工作项状态流转则可以通过获取状态方案中的工作项状态流转列表获取。

assignee_id可选	String	
工作项负责人的id。

parent_id可选	String	
工作项的父工作项的id。当前工作项类型支持的父工作类型包含parent_id对应的工作项类型时，该参数有效。具体配置见属性与视图子工作项配置。

sprint_id可选	String	
所属迭代id。该字段只有项目类型为scrum和hybrid时有效。

version_id可选	String	
所属发布的id。

board_id可选	String	
看板的id。该字段只有项目类型为kanban和hybrid时有效。

entry_id可选	String	
看板栏的id。该字段只有项目类型为kanban和hybrid时有效。

swimlane_id可选	String	
泳道的id。该字段只有项目类型为kanban和hybrid时有效。

story_points可选	Number	
工作项的故事点。当工作项的属性在视图中配置了故事点属性时，该参数生效。故事点数值必须是大于等于0的整数或最多包含一位小数的正数。

estimated_workload可选	Number	
工作项的预估工时。

remaining_workload可选	Number	
工作项的剩余工时。

properties可选	Object	
工作项属性的键值对集合，需要注意的是，当前工作项类型对应的工作项属性方案需要包含这些工作项属性，例如工作项属性方案中包含prop_a和prop_b。

properties.prop_a可选	Object	
工作项属性prop_a。

properties.prop_b可选	Object	
工作项属性prop_b。

participant_ids可选	String[]	
工作项关注人的id列表。

请求示例：
{
    "project_id": "5eb623f6a70571487ea47000",
    "type_id": "bug",
    "title": "这是一个缺陷",
    "description": "这是一个缺陷的描述",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "state_id": "5c9b35de90ad7153c2062f18",
    "parent_id": "5edca112b06305c524cad2fa",
    "sprint_id": "5ecf7b74eaab845a2aa53138",
    "version_id": "5eb623f6a70571487ea47001",
    "board_id": "5eb623f6a70571487ea47222",
    "entry_id": "5ee1c4fac5b3c51206f0a861",
    "swimlane_id": "5ee1c4fac5b3c51206f0a866",
    "priority_id": "5eb623f6a70571487ea47111",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "participant_ids": [
        "a0417f68e846aae315c85d24643678a9"
    ],
    "story_points": 1,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "5edca524cad2fa1125cb0630",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SCR-5",
    "title": "这是一个缺陷",
    "type": "bug",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": "5edca112b06305c524cad2fa",
    "short_id": "1bAqLmTG",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/1bAqLmTG",
    "parent": {
        "id": "5edca112b06305c524cad2fa",
        "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
        "identifier": "SCR-3",
        "title": "这是一个用户故事",
        "type": "story",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06309c",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value",
            "risk": null,
            "business_value": null,
            "effort": null,
            "backlog_type": null,
            "backlog_from": null
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
        "name": "1.0.1",
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
    "sprint": {
        "id": "5ecf7b74eaab845a2aa53138",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
        "name": "Sprint 1",
        "start_at": 1589791860,
        "end_at": 1589791860,
        "status": "completed"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board": null,
    "entry": null,
    "swimlane": null,
    "phase": null,
    "description": "这是一个缺陷的描述",
    "completed_at": 1583290347,
    "story_points": 1,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "severity": null,
        "replay_version": null,
        "reappear_probability": null,
        "bug_type": null,
        "reason": null,
        "solution": null,
        "replay_step": null
    },
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
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
部分更新一个工作项
PATCH
https://rest_api_root/v1/project/work_items/{work_item_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

请求参数
字段	类型	描述
title可选	String	
工作项的标题。

description可选	String	
工作项的描述。

start_at可选	Number	
工作项的开始时间。当工作项类型为里程碑时，该参数无效。

end_at可选	Number	
工作项的截止时间。

priority_id可选	String	
工作项优先级的id。

state_id可选	String	
工作项状态的id。工作项状态的id在设置时必须同时满足工作项类型的工作项状态方案和工作项状态流转的状态值才能成功完成设置。工作项状态方案可以通过获取工作项状态方案列表获取。工作项状态流转则可以通过获取状态方案中的工作项状态流转列表获取。

assignee_id可选	String	
工作项负责人的id。

parent_id可选	String	
工作项的父工作项的id。当前工作项类型支持的父工作类型包含parent_id对应的工作项类型时，该参数有效。具体配置见属性与视图子工作项配置。

version_id可选	String	
所属发布的id。

board_id可选	String	
看板的id。该字段只有项目类型为kanban和hybrid时有效。

entry_id可选	String	
看板栏的id。该字段只有项目类型为kanban和hybrid时有效。

swimlane_id可选	String	
泳道的id。该字段只有项目类型为kanban和hybrid时有效。

phase_id可选	String	
所属计划的id。该字段只有项目类型为waterfall和hybrid时有效。

story_points可选	Number	
工作项的故事点。当工作项的属性在视图中配置了故事点属性时，该参数生效。故事点数值必须是大于等于0的整数或最多包含一位小数的正数。

estimated_workload可选	Number	
工作项的预估工时。

remaining_workload可选	Number	
工作项的剩余工时。

properties可选	Object	
工作项属性的键值对集合，需要注意的是，当前工作项类型对应的工作项属性方案需要包含这些工作项属性，例如工作项属性方案中包含prop_a和prop_b。

properties.prop_a可选	Object	
工作项属性prop_a。

properties.prop_b可选	Object	
工作项属性prop_b。

请求示例：
{
    "title": "这是一个用户故事",
    "description": "这是一个用户故事的描述",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "sprint_id": "5ecf7b74eaab845a2aa53138",
    "version_id": "5eb623f6a70571487ea47001",
    "priority_id": "5eb623f6a70571487ea47111",
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "story_points": 1,
    "state_id": "5c9b35de90ad7153c2062f18",
    "parent_id": "5edfa3b67463c1ee626c0979",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
响应示例：
{
    "id": "5edca112b06305c524cad2fa",
    "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SCR-3",
    "title": "这是一个用户故事",
    "type": "story",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": "5edfa3b67463c1ee626c0979",
    "short_id": "b9WqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/b9WqLmTO",
    "parent": {
        "id": "5edfa3b67463c1ee626c0979",
        "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0979",
        "identifier": "SCR-2",
        "title": "这是一个特性",
        "type": "feature",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06306c",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value",
            "risk": null,
            "business_value": null,
            "effort": null,
            "backlog_type": null,
            "backlog_from": null
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
        "name": "1.0.1",
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
    "sprint": {
        "id": "5ecf7b74eaab845a2aa53138",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
        "name": "Sprint 1",
        "start_at": 1589791860,
        "end_at": 1589791860,
        "status": "completed"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board": null,
    "entry": null,
    "swimlane": null,
    "phase": null,
    "story_points": 1,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "description": "这是一个用户故事的描述",
    "completed_at": 1583290347,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "risk": null,
        "backlog_type": null,
        "backlog_from": null
    },
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
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
批量部分更新工作项属性
PATCH
https://rest_api_root/v1/project/work_items
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
ids	String[]	
需要更新的工作项的id列表。最多可以批量更新100个工作项。

property_name	String	
需要更新的工作项属性名。

允许值: title, start_at, end_at, description, priority_id, assignee_id, state_id, story_points, estimated_workload, remaining_workload, 自定义属性

property_value可选	String	
需要更新的工作项属性值。

请求示例：
{
    "ids": [
        "547000eb6a70571487623fea"
    ],
    "property_name": "title",
    "property_value": "这是一个工作项"
}
响应示例：
{
    "inserts": 0,
    "updates": 1,
    "deletes": 0
}
获取工作项列表
GET
https://rest_api_root/v1/project/work_items
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
identifier可选	String	
工作项编号。

project_ids可选	String	
项目的id，使用','分割，最多只能20个。

type_ids可选	String	
工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。最多只能20个。

parent_ids可选	String	
父工作项的id，使用','分割，最多只能20个。

assignee_ids可选	String	
工作项负责人的id，使用','分割，最多只能20个。

state_ids可选	String	
工作项状态的id，使用','分割，最多只能20个。

start_between可选	String	
开始时间介于的时间范围，通过','分割起始时间。比如1580000000,1590000000表示开始时间介于两个时间之间；,1590000000表示开始时间小于该时间；1580000000表示开始时间大于该时间。

end_between可选	String	
结束时间介于的时间范围，通过','分割起始时间。使用方式参考开始时间。

priority_ids可选	String	
工作项优先级的id，使用','分割，最多只能20个。

bug_type_ids可选	String	
缺陷类别的id，使用','分割，最多只能20个。

tag_ids可选	String	
工作项标签的id，使用','分割，最多只能20个。

sprint_ids可选	String	
迭代的id，使用','分割，最多只能20个。

board_ids可选	String	
看板的id，使用','分割，最多只能20个。

entry_ids可选	String	
看板栏的id，使用','分割，最多只能20个。

swimlane_ids可选	String	
泳道的id，使用','分割，最多只能20个。

phase_ids可选	String	
所属计划的id，使用','分割，最多只能20个。

version_ids可选	String	
发布的id，使用','分割，最多只能20个。

created_by_ids可选	String	
创建人的id，使用','分割，最多只能20个。

created_between可选	String	
创建时间介于的时间范围，通过','分割起始时间。使用方式参考开始时间。

updated_between可选	String	
更新时间介于的时间范围，通过','分割起始时间。使用方式参考开始时间。

participant_id可选	String	
工作项关注人的id。

keywords可选	String	
关键字。支持工作项编号和工作项标题。

include_public_image_token可选	String	
包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.prop_b。

include_deleted可选	Boolean	
是否查询已删除的工作项。该值默认为false。

include_archived可选	Boolean	
是否查询已归档的工作项。该值默认为false。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 4,
    "values": [
        {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa112b06305c",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SCR-1",
            "title": "这是一个史诗",
            "type": "epic",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": null,
            "short_id": "d9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/d9WqLmTO",
            "parent": null,
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "version": null,
            "sprint": null,
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": null,
            "description": "这是一个史诗的描述",
            "completed_at": 1583290347,
            "story_points": null,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "business_value": null,
                "effort": 123,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca524cad2fa112b06305c",
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
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Df",
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
        },
        {
            "id": "5edfa3b67463c1ee626c0979",
            "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0979",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SCR-2",
            "title": "这是一个特性",
            "type": "feature",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06305c",
            "short_id": "a9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
            "parent": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa112b06305c",
                "identifier": "SCR-1",
                "title": "这是一个史诗",
                "type": "epic",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": null,
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value",
                    "risk": null,
                    "business_value": null,
                    "effort": null,
                    "backlog_type": null,
                    "backlog_from": null
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "version": {
                "id": "5eb623487ea47001f6a70571",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
                "name": "1.0.1",
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
            "sprint": null,
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": null,
            "description": "这是一个特性的描述",
            "completed_at": 1583290347,
            "story_points": null,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "business_value": null,
                "effort": null,
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edfa3b67463c1ee626c0979",
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
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-fC",
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
            "is_archived": 1,
            "is_deleted": 1
        },
        {
            "id": "5edca112b06305c524cad2fa",
            "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SCR-3",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edfa3b67463c1ee626c0979",
            "short_id": "b9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/b9WqLmTO",
            "parent": {
                "id": "5edfa3b67463c1ee626c0979",
                "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0979",
                "identifier": "SCR-2",
                "title": "这是一个特性",
                "type": "feature",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa112b06305g",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value",
                    "risk": null,
                    "business_value": null,
                    "effort": null,
                    "backlog_type": null,
                    "backlog_from": null
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "version": {
                "id": "5eb623487ea47001f6a70571",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
                "name": "1.0.1",
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
            "sprint": {
                "id": "5ecf7b74eaab845a2aa53138",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
                "name": "Sprint 1",
                "start_at": 1589791860,
                "end_at": 1589791860,
                "status": "completed"
            },
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": null,
            "description": "这是一个用户故事的描述",
            "completed_at": 1583290347,
            "story_points": 1,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
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
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Hm",
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
        },
        {
            "id": "5edca5d2fa112b06305c24ca",
            "url": "https://rest_api_root/v1/project/work_items/5edca5d2fa112b06305c24ca",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "WTF-2",
            "title": "这是一个瀑布项目下需求类型的工作项",
            "type": "630da48bc9443b1aa94ce3ea",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "version": null,
            "sprint": null,
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": {
                "id": "63761fee31caaf77189816b4",
                "url": "http://rest_api_root/v1/project/projects/63761fee31caaf7718981698/phases/63761fee31caaf77189816b4",
                "title": "这是一个阶段",
                "identifier": "WTF-1"
            },
            "description": "这是一个瀑布项目下需求类型的工作项描述",
            "completed_at": 1583290347,
            "story_points": null,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
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
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Ki",
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
        },
        {
            "id": "6efca131b06329c524cdd2fb",
            "url": "https://rest_api_root/v1/project/work_items/6efca131b06329c524cdd2fb",
            "project": {
                "id": "5eb623f6a70571487ea47004",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004",
                "identifier": "HBR",
                "name": "Hybrid项目",
                "type": "hybrid",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "HBR-1",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edfa3b67463c1ee626c0980",
            "short_id": "e9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/e9WqLmTO",
            "parent": {
                "id": "5edfa3b67463c1ee626c0980",
                "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0980",
                "identifier": "HBR-2",
                "title": "这是一个特性",
                "type": "feature",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": null,
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value",
                    "risk": null,
                    "business_value": null,
                    "effort": null,
                    "backlog_type": null,
                    "backlog_from": null
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "version": {
                "id": "5eb623487ea47001f6a70571",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/versions/5eb623487ea47001f6a70582",
                "name": "1.0.1",
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
            "sprint": {
                "id": "5ecf7b74eaab845a2aa53139",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/sprints/5ecf7b74eaab845a2aa53139",
                "name": "Sprint 1",
                "start_at": 1589791860,
                "end_at": 1589791860,
                "status": "completed"
            },
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": {
                "id": "5eb623f6a70571487ea47333",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333",
                "name": "kanban",
                "work_item_types": [
                    "epic",
                    "feature",
                    "issue",
                    "story"
                ]
            },
            "entry": {
                "id": "5ee1c4fac5b3c51206f0a862",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/entries/5ee1c4fac5b3c51206f0a862",
                "name": "需求池"
            },
            "swimlane": {
                "id": "5ee1c4fac5b3c51206f0a867",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/swimlanes/5ee1c4fac5b3c51206f0a867",
                "name": "默认泳道"
            },
            "phase": {
                "id": "63761fee31caaf77189816b5",
                "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/phases/63761fee31caaf77189816b5",
                "title": "这是一个阶段",
                "identifier": "WTF-1"
            },
            "description": "<p>这是一个用户故事的描述<p><img src=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" originUrl=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" alt=\"图片.png\" size=\"35460\" style=\"text-align: center;\" />",
            "completed_at": 1583290347,
            "story_points": 1,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
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
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Ac",
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
删除一个工作项
DELETE
https://rest_api_root/v1/project/work_items/{work_item_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

响应示例：
{
    "id": "5edca5d2fa112b06305c24ca",
    "url": "https://rest_api_root/v1/project/work_items/5edca5d2fa112b06305c24ca",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KB",
        "name": "看板项目",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "KB-1",
    "title": "这是一个事务",
    "type": "issue",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "short_id": "c9WqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "version": null,
    "sprint": null,
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board": {
        "id": "5eb623f6a70571487ea47222",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
        "name": "kanban",
        "work_item_types": [
            "epic",
            "feature",
            "issue",
            "story"
        ]
    },
    "entry": {
        "id": "5ee1c4fac5b3c51206f0a861",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222/entries/5ee1c4fac5b3c51206f0a861",
        "name": "需求池"
    },
    "swimlane": {
        "id": "5ee1c4fac5b3c51206f0a866",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222/swimlanes/5ee1c4fac5b3c51206f0a866",
        "name": "默认泳道"
    },
    "phase": null,
    "description": "这是一个事务的描述",
    "completed_at": 1583290347,
    "story_points": null,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "entry_status": null,
        "entry_position": null,
        "operation_time": 1691571221
    },
    "tags": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
            "name": "标签-1"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca5d2fa112b06305c24ca",
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
    "is_deleted": 1
}
获取工作项类型列表
GET
https://rest_api_root/v1/project/work_item/types?project_id={project_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id	String	
项目的id。

响应示例（scrum项目）：
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement"
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement"
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        }
    ]
}
响应示例（kanban项目）：
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement"
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement"
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "issue",
            "url": "http://rest_api_root/v1/project/work_item_types/issue",
            "name": "事务",
            "group": "issue"
        }
    ]
}
响应示例（waterfall项目）：
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
            "name": "需求",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "https://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "https://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ee",
            "name": "阶段",
            "group": "plan"
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ef",
            "name": "里程碑",
            "group": "plan"
        }
    ]
}
响应示例（hybrid项目）：
{
    "page_size": 30,
    "page_index": 0,
    "total": 9,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement"
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement"
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "issue",
            "url": "http://rest_api_root/v1/project/work_item_types/issue",
            "name": "事务",
            "group": "issue"
        },
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
            "name": "需求",
            "group": "requirement"
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ee",
            "name": "阶段",
            "group": "plan"
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ef",
            "name": "里程碑",
            "group": "plan"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        }
    ]
}
获取工作项状态列表
GET
https://rest_api_root/v1/project/work_item/states?project_id={project_id}&work_item_type_id={work_item_type_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id	String	
项目的id。

work_item_type_id	String	
工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b35de90ad7153c2062f18",
            "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
            "name": "新建",
            "type": "pending",
            "color": "#ff7575"
        }
    ]
}
获取工作项属性列表
GET
https://rest_api_root/v1/project/work_item/properties?project_id={project_id}&work_item_type_id={work_item_type_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id	String	
项目的id。

work_item_type_id	String	
工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "severity",
            "url": "https://rest_api_root/v1/project/work_item/work_item_properties/severity",
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
            ]
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/project/work_item_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null
        }
    ]
}
获取工作项优先级列表
GET
https://rest_api_root/v1/project/work_item/priorities?project_id={project_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47111",
            "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
            "name": "最高"
        }
    ]
}
获取工作项标签列表
GET
https://rest_api_root/v1/project/work_item/tags
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id	String	
项目的id。

name可选	String	
标签的名称。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
            "name": "标签-1"
        }
    ]
}
向工作项中添加一个标签
POST
https://rest_api_root/v1/project/work_items/{work_item_id}/tags
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

请求参数
字段	类型	描述
tag_id	String	
标签的id。

请求示例：
{
    "tag_id": "5e6b35de50ef8153c2062f70"
}
响应示例：
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630/tags/5e6b35de50ef8153c2062f70",
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b05105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "tag": {
        "id": "5e6b35de50ef8153c2062f70",
        "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
        "name": "标签-1"
    }
}
在工作项中移除一个标签
DELETE
https://rest_api_root/v1/project/work_items/{work_item_id}/tags/{tag_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

tag_id	String	
标签的id。

响应示例：
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630/tags/5e6b35de50ef8153c2062f70",
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b05105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "tag": {
        "id": "5e6b35de50ef8153c2062f70",
        "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
        "name": "标签-1"
    }
}
获取工作项关联类型列表
GET
https://rest_api_root/v1/project/work_item/relation_types
权限: 企业令牌/用户令牌

响应示例：
{
    "page_index": 0,
    "page_size": 30,
    "total": 1,
    "values": [
        {
            "id": "676510af06fd48a4a4e12616",
            "url": "https://rest_api_root/v1/project/work_item/relation_types/676510af06fd48a4a4e12616",
            "name": "重复",
            "category": "duplicate",
            "is_system": 1
        }
    ]
}
关联一个工作项
POST
https://rest_api_root/v1/project/work_items/{work_item_id}/relations
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

请求参数
字段	类型	描述
target_work_item_id	String	
目标工作项的id。

relation_type	String	
关联的类型。

允许值: mention, clone, cloned_by, duplicate, relate, cause, caused_by, block, blocked_by, dependency, 自定义关联类型的id

请求示例：
{
    "target_work_item_id": "5f9a65ef20ef8153c1462e64",
    "relation_type": "relate"
}
响应示例：
{
    "id": "58fb35de50ef8153c2062e36",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112/relations/58fb35de50ef8153c2062e36",
    "relation_type": "relate",
    "origin_work_item": {
        "id": "5edca524cad2b06305cfa112",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112",
        "identifier": "SCR-4",
        "title": "这是一个任务",
        "type": "task",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_work_item": {
        "id": "5f9a65ef20ef8153c1462e64",
        "url": "https://rest_api_root/v1/project/work_items/5f9a65ef20ef8153c1462e64",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": null,
        "short_id": "a9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
获取关联的工作项列表
GET
https://rest_api_root/v1/project/work_items/{work_item_id}/relations
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

查询参数
字段	类型	描述
relation_type可选	String	
关联的类型。

允许值: mention, clone, cloned_by, duplicate, relate, cause, caused_by, block, blocked_by, dependency, 自定义关联类型的id

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "58fb35de50ef8153c2062e36",
            "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112/relations/58fb35de50ef8153c2062e36",
            "relation_type": "relate",
            "origin_work_item": {
                "id": "5edca524cad2b06305cfa112",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112",
                "identifier": "SCR-4",
                "title": "这是一个任务",
                "type": "task",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa1125cb0629",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "target_work_item": {
                "id": "5f9a65ef20ef8153c1462e64",
                "url": "https://rest_api_root/v1/project/work_items/5f9a65ef20ef8153c1462e64",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": null,
                "short_id": "a9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            }
        }
    ]
}
取消关联一个工作项
DELETE
https://rest_api_root/v1/project/work_items/{work_item_id}/relations/{relation_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

relation_id	String	
工作项关联的id。

响应示例：
{
    "id": "58fb35de50ef8153c2062e36",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112/relations/58fb35de50ef8153c2062e36",
    "relation_type": "relate",
    "origin_work_item": {
        "id": "5edca524cad2b06305cfa112",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112",
        "identifier": "SCR-4",
        "title": "这是一个任务",
        "type": "task",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_work_item": {
        "id": "5f9a65ef20ef8153c1462e64",
        "url": "https://rest_api_root/v1/project/work_items/5f9a65ef20ef8153c1462e64",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": null,
        "short_id": "a9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
获取工作项流转记录列表
GET
https://rest_api_root/v1/project/work_items/{work_item_id}/transition_histories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_id	String	
工作项的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630/transition_histories/5e6b35de50ef8153c2062f70",
            "work_item": {
                "id": "5edca524cad2fa1125cb0630",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1674493200,
                "end_at": 1674493200,
                "parent_id": "5edca524cad2fa112b05105c",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "from_state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#56ABFB"
            },
            "to_state": {
                "id": "5ef85b1e9481936604da7f4c",
                "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
                "name": "进行中",
                "type": "in_progress",
                "color": "#F6C659"
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
创建一个工作项交付目标
POST
https://rest_api_root/v1/project/deliverables
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
work_item_id	String	
工作项的id。工作项所属项目类型必须为waterfall或hybrid。

name	String	
工作项交付目标的名称。

content_type可选	String	
工作项交付物的类型。工作项交付物的类型。只支持link。attachment类型的交付物通过向交付目标中上传一个文件接口上传附件。

content可选	String	
工作项交付物。当工作项交付物的类型是link时，工作项交付物为包含name和href的对象，例如：{ "name": "链接工作项交付目标", "href": "https://rest_api_root/wiki/spaces/public/pages/6472e6f2f1968d3fdb0aba15" }。当工作项交付物不为空时，参数必须包含交付物类型。

请求示例：
{
    "work_item_id": "63761fee31caaf77189816b4",
    "name": "工作项交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    }
}
响应示例：
{
    "id": "63761fee31caaf7718981876",
    "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
    "name": "阶段交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    },
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "63761fee31caaf77189816b4",
        "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
        "identifier": "WTF-5",
        "title": "这是一个阶段",
        "type": "630da48bc9443b1aa94ce3ee",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "created_at": 1668685806,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1668685806,
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
部分更新一个工作项交付目标
PATCH
https://rest_api_root/v1/project/deliverables/{deliverable_target_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
deliverable_target_id	String	
工作项交付目标id。

请求参数
字段	类型	描述
work_item_id可选	String	
工作项的id。

name可选	String	
工作项交付目标的名称。

content_type可选	String	
工作项交付物的类型。只支持link。attachment类型的交付物通过向交付目标中上传一个文件接口上传附件。

content可选	String	
工作项交付物。当工作项交付物的类型是link时，工作项交付物为包含name和href的对象，例如：{ "name": "链接工作项交付目标", "href": "https://rest_api_root/wiki/spaces/public/pages/6472e6f2f1968d3fdb0aba15" }。当工作项交付物不为空时，参数必须包含交付物类型。

请求示例：
{
    "work_item_id": "63761fee31caaf77189816b4",
    "name": "工作项交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    }
}
响应示例：
{
    "id": "63761fee31caaf7718981876",
    "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
    "name": "阶段交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    },
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "63761fee31caaf77189816b4",
        "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
        "identifier": "WTF-5",
        "title": "这是一个阶段",
        "type": "630da48bc9443b1aa94ce3ee",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "created_at": 1668685806,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1668687806,
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
获取工作项交付目标列表
GET
https://rest_api_root/v1/project/deliverables
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id可选	String	
项目的id。项目类型必须为waterfall或hybrid。

work_item_id可选	String	
工作项的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "63761fee31caaf7718981876",
            "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
            "name": "阶段交付目标",
            "content_type": "link",
            "content": {
                "name": "PingCode",
                "href": "https://www.pingcode.com"
            },
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item": {
                "id": "63761fee31caaf77189816b4",
                "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
                "identifier": "WTF-5",
                "title": "这是一个阶段",
                "type": "630da48bc9443b1aa94ce3ee",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa1125cb0629",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "created_at": 1668685806,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1668685806,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "63761fee31caaf7718981877",
            "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981877",
            "name": "缺陷交付目标",
            "content_type": "attachment",
            "content": {
                "id": "64abd9050461799795b6ea3e",
                "url": "https://rest_api_root/v1/attachments/64abd9050461799795b6ea3e?deliverable_target_id=63761fee31caaf7718981877",
                "title": "fixed.png",
                "size": 11396,
                "type": "file"
            },
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item": {
                "id": "63761fee31caaf77189816b5",
                "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b5",
                "identifier": "WTF-6",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290319,
                "end_at": 1583290357,
                "parent_id": "5edca524cad2fa1125cb0623",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "created_at": 1668685816,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1668685816,
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
删除一个工作项交付目标
DELETE
https://rest_api_root/v1/project/deliverables/{deliverable_target_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
deliverable_target_id	String	
工作项交付目标的id。

响应示例：
{
    "id": "63761fee31caaf7718981876",
    "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
    "name": "阶段交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    },
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "63761fee31caaf77189816b4",
        "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
        "identifier": "WTF-5",
        "title": "这是一个阶段",
        "type": "630da48bc9443b1aa94ce3ee",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "created_at": 1668685806,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1668685806,
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
发布
用于查看和管理发布相关信息。
资源地址：{GET} https://rest_api_root/v1/project/projects/{project_id}/versions/{version_id}

资源属性
字段	类型	描述
id	String	
发布的id。

url	String	
发布的url。

project	Object	
发布所属的项目。

name	String	
发布的名称。

start_at	Number	
发布的开始时间。

end_at	Number	
发布的结束时间。

stage	Object	
发布的当前阶段。

assignee	Object	
发布的负责人。

stages	Object[]	
发布的阶段列表。

progress	Number	
发布的进度。

changelog	String	
发布的发布日志。

operate_at	Number	
发布的最后操作时间。

categories	Object[]	
发布的类别列表。

created_at	Number	
发布的创建时间。

created_by	Object	
发布的创建人。

updated_at	Number	
发布的更新时间。

updated_by	Object	
发布的更新人。

全量数据示例：
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": 1565366400
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": "发布日志",
    "operate_at": 1565366400,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
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
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    }
}
创建一个发布
POST
https://rest_api_root/v1/project/projects/{project_id}/versions
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
发布的名称。同一项目下该名称是唯一的。

start_at	Number	
开始的时间。

end_at	Number	
发布的时间。

assignee_id	String	
负责人的id。

stage_id可选	String	
发布阶段的id。

category_ids可选	String[]	
发布类别id数组。

请求示例：
{
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "stage_id": "5f44a8f8bb347b14b49624a1",
    "category_ids": [
        "676a460a0fd987b7ea320889"
    ]
}
响应示例：
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": null,
    "operate_at": 1565255712,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
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
批量创建发布
POST
https://rest_api_root/v1/project/versions/bulk
权限: 企业令牌

请求参数
字段	类型	描述
versions	Object[]	
需要批量创建的发布。该参数是一个对象数组（数组内对象不得超过100个）。

versions.project_id	String	
发布所属项目的id。

versions.name	String	
发布的名称。

versions.start_at	Number	
发布的开始时间。

versions.end_at	Number	
发布的时间。

versions.assignee_id	String	
发布负责人的id。

versions.stage_id	String	
发布的阶段id。

versions.category_ids可选	String[]	
发布类别的id列表。

请求示例：
{
    "versions": [
        {
            "project_id": "5eb623f6a70571487ea47000",
            "name": "1.0.0",
            "start_at": 1565193600,
            "end_at": 1566403200,
            "assignee_id": "a0417f68e846aae315c85d24643678a9",
            "stage_id": "5f44a8f8bb347b14b49624a1",
            "category_ids": [
                "676a460a0fd987b7ea320889"
            ]
        }
    ]
}
响应示例：
[
    {
        "state": "success",
        "version": {
            "id": "5eb623f6a70571487ea47001",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "type": "scrum",
                "name": "Scrum项目",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "1.0.0",
            "start_at": 1565193600,
            "end_at": 1566403200,
            "stage": {
                "id": "5f44a8f8bb347b14b49624a1",
                "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                "name": "未开始",
                "type": "pending",
                "color": "#FA8888"
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "stages": [
                {
                    "id": "5f44a8f8bb347b14b49624a1",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                    "name": "未开始",
                    "operate_at": 1565366400
                },
                {
                    "id": "5f44a8f8bb347b14b49624a2",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
                    "name": "进行中",
                    "operate_at": null
                },
                {
                    "id": "5f44a8f8bb347b14b49624a3",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
                    "name": "已发布",
                    "operate_at": null
                }
            ],
            "progress": 0,
            "changelog": null,
            "operate_at": 1565366400,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320889",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
                    "name": "私有部署发布"
                }
            ],
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
    }
]
部分更新一个发布
PATCH
https://rest_api_root/v1/project/projects/{project_id}/versions/{version_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

version_id	String	
发布的id。

请求参数
字段	类型	描述
name可选	String	
发布的名称。同一项目下该名称是唯一的。

start_at可选	Number	
开始的时间。

end_at可选	Number	
发布的时间。

assignee_id可选	String	
负责人的id。

stage_id可选	String	
发布阶段的id。

operate_at可选	Number	
发布阶段的日期。需要和stage_id一起传递。

category_ids可选	String[]	
发布类别id数组。

请求示例：
{
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "assignee_id": "a0417f68e846aae315c85d24643678a9",
    "stage_id": "5f44a8f8bb347b14b49624a1",
    "operate_at": 1565366400,
    "category_ids": [
        "676a460a0fd987b7ea320889"
    ]
}
响应示例：
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": 1565366400
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": "发布日志",
    "operate_at": 1565366400,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
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
获取发布列表
GET
https://rest_api_root/v1/project/projects/{project_id}/versions
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

查询参数
字段	类型	描述
name可选	String	
发布的名字。

status可选	String	
发布的状态。

允许值: pending, in_progress, published

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
            "id": "5eb623f6a70571487ea47001",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "type": "scrum",
                "name": "Scrum项目",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "1.0.0",
            "start_at": 1565193600,
            "end_at": 1566403200,
            "stage": {
                "id": "5f44a8f8bb347b14b49624a1",
                "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                "name": "未开始",
                "type": "pending",
                "color": "#FA8888"
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "stages": [
                {
                    "id": "5f44a8f8bb347b14b49624a1",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                    "name": "未开始",
                    "operate_at": 1565366400
                },
                {
                    "id": "5f44a8f8bb347b14b49624a2",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
                    "name": "进行中",
                    "operate_at": null
                },
                {
                    "id": "5f44a8f8bb347b14b49624a3",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
                    "name": "已发布",
                    "operate_at": null
                }
            ],
            "progress": 0,
            "changelog": "发布日志",
            "operate_at": 1565366400,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320889",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
                    "name": "私有部署发布"
                }
            ],
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
删除一个发布
DELETE
https://rest_api_root/v1/project/projects/{project_id}/versions/{version_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

version_id	String	
发布的id。

响应示例：
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": null,
    "operate_at": 1565255712,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
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
创建一个发布阶段
POST
https://rest_api_root/v1/project/stages
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
发布阶段的名称。在一个企业中这个名称是唯一的。

type	String	
发布阶段的类型。

允许值: pending, in_progress, published

请求示例：
{
    "name": "新建",
    "type": "pending"
}
响应示例：
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "pending",
    "color": "#ff7575"
}
部分更新一个发布阶段
PATCH
https://rest_api_root/v1/project/stages/{stage_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
stage_id	String	
发布阶段的id。

请求参数
字段	类型	描述
name可选	String	
发布阶段的名称。在一个企业中这个名称是唯一的。

type可选	String	
发布阶段的类型。

允许值: pending, in_progress, published

请求示例：
{
    "name": "新建",
    "type": "in_progress"
}
响应示例：
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "in_progress",
    "color": "#ff7575"
}
获取发布阶段列表
GET
https://rest_api_root/v1/project/stages
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b35de90ad7153c2062f18",
            "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
            "name": "新建",
            "type": "in_progress",
            "color": "#ff7575"
        }
    ]
}
删除一个发布阶段
DELETE
https://rest_api_root/v1/project/stages/{stage_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
stage_id	String	
发布阶段的id。

请求参数
字段	类型	描述
replace_id可选	String	
替换的发布阶段id，如果一个发布阶段已经被发布使用，删除该发布阶段时需要提供一个替换的发布阶段。

请求示例：
{
    "replace_id": "5c9b35de90ad7153c2062f19"
}
响应示例：
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "in_progress",
    "color": "#ff7575"
}
创建一个发布分组
POST
https://rest_api_root/v1/project/projects/{project_id}/version_sections
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
发布分组的名称。

description可选	String	
发布分组的描述。

请求示例：
{
    "name": "私有部署",
    "description": "私有部署发布分组"
}
响应示例：
{
    "id": "63560f3ad02cbc4f9df91251",
    "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署",
    "description": "私有部署发布分组"
}
部分更新一个发布分组
PATCH
https://rest_api_root/v1/project/projects/{project_id}/version_sections/{section_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

section_id	String	
发布分组的id。

请求参数
字段	类型	描述
name	String	
发布分组的名称。

description可选	String	
发布分组的描述。

请求示例：
{
    "name": "私有部署",
    "description": "私有部署发布分组"
}
响应示例：
{
    "id": "63560f3ad02cbc4f9df91251",
    "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署",
    "description": "私有部署发布分组"
}
获取发布分组列表
GET
https://rest_api_root/v1/project/projects/{project_id}/version_sections
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63560f3ad02cbc4f9df91251",
            "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "这是一个发布分组",
            "description": "这是一个发布分组的描述"
        }
    ]
}
删除一个发布分组
DELETE
https://rest_api_root/v1/project/projects/{project_id}/version_sections/{section_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

section_id	String	
发布分组的id。

响应示例：
{
    "id": "63560f3ad02cbc4f9df91252",
    "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91252",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "临时部署",
    "description": "临时发布分组"
}
创建一个发布类别
POST
https://rest_api_root/v1/project/projects/{project_id}/version_categories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

请求参数
字段	类型	描述
name	String	
发布类别的名称。

section_id可选	String	
发布类别所属发布分组。

请求示例：
{
    "name": "私有部署发布",
    "section_id": "63560f3ad02cbc4f9df91251"
}
响应示例：
{
    "id": "676a460a0fd987b7ea320889",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署发布",
    "section": {
        "id": "63560f3ad02cbc4f9df91251",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
        "name": "私有部署发布分组"
    }
}
部分更新一个发布类别
PATCH
https://rest_api_root/v1/project/projects/{project_id}/version_categories/{version_category_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

version_category_id	String	
发布类别的id。

请求参数
字段	类型	描述
name可选	String	
发布类别的名称。

section_id可选	String	
发布类别所属发布分组。

请求示例：
{
    "name": "私有部署发布",
    "section_id": "63560f3ad02cbc4f9df91251"
}
响应示例：
{
    "id": "676a460a0fd987b7ea320889",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署发布",
    "section": {
        "id": "63560f3ad02cbc4f9df91251",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
        "name": "私有部署发布分组"
    }
}
获取发布类别列表
GET
https://rest_api_root/v1/project/projects/{project_id}/version_categories
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "私有部署发布",
            "section": {
                "id": "63560f3ad02cbc4f9df91251",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
                "name": "私有部署发布分组"
            }
        }
    ]
}
删除一个发布类别
DELETE
https://rest_api_root/v1/project/projects/{project_id}/version_categories/{version_category_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
project_id	String	
项目的id。

version_category_id	String	
发布类别的id。

响应示例：
{
    "id": "676a460a0fd987b7ea320890",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320890",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "这是一个发布类别",
    "section": {
        "id": "63560f3ad02cbc4f9df91251",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
        "name": "私有部署发布分组"
    }
}
项目配置中心
项目配置
用于查看和管理项目相关配置。

创建一个项目属性
POST
https://rest_api_root/v1/project/project_properties
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
项目属性的名称。在一个企业中这个名称是唯一的。

type	String	
项目属性的类型。

允许值: text, textarea, select, multi_select, cascade_select, cascade_multi_select, member, members, date, number, progress, rate, link

options可选	Object[]	
选择项列表。当属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。

options._id可选	String	
选择项id。该值在选择项中是唯一的。

options.text	String	
选择项内容。该值在选择项中是唯一的。

options.parent_id可选	String	
选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。

请求示例：
{
    "name": "项目风险",
    "type": "select",
    "options": [
        {
            "text": "高"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "中"
        },
        {
            "_id": "5efb1859110533727a82c605",
            "text": "低"
        }
    ]
}
响应示例：
{
    "id": "risk",
    "url": "https://rest_api_root/v1/project/project_properties/risk",
    "name": "项目风险",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "高"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "中"
        },
        {
            "_id": "5efb1859110533727a82c605",
            "text": "低"
        }
    ],
    "is_removable": false,
    "is_name_editable": false,
    "is_options_editable": false
}
获取项目属性列表
GET
https://rest_api_root/v1/project/project_properties
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "risk",
            "url": "https://rest_api_root/v1/project/project_properties/risk",
            "name": "项目风险",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "高"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "中"
                },
                {
                    "_id": "5efb1859110533727a82c605",
                    "text": "低"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "name",
            "url": "https://rest_api_root/v1/project/project_properties/name",
            "name": "名称",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
工作项配置
用于查看和管理工作项相关配置。

创建一个工作项标签
POST
https://rest_api_root/v1/project/work_item_tags
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
标签的名称。在一个企业中这个名称是唯一的。

请求示例：
{
    "name": "标签-1"
}
响应示例：
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
    "name": "标签-1",
    "color": "#56ABFB"
}
部分更新一个工作项标签
PATCH
https://rest_api_root/v1/project/work_item_tags/{tag_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
tag_id	String	
标签的id。

请求参数
字段	类型	描述
name可选	String	
标签的名称。在一个企业中这个名称是唯一的。

请求示例：
{
    "name": "标签-2"
}
响应示例：
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
    "name": "标签-2",
    "color": "#56ABFB"
}
获取全部工作项标签列表
GET
https://rest_api_root/v1/project/work_item_tags
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
name可选	String	
标签的名称。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
            "name": "标签-2",
            "color": "#56ABFB"
        }
    ]
}
删除一个工作项标签
DELETE
https://rest_api_root/v1/project/work_item_tags/{tag_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
tag_id	String	
标签的id。

响应示例：
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
    "name": "标签-2",
    "color": "#56ABFB"
}
创建一个工作项类型
POST
https://rest_api_root/v1/project/work_item_types
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
工作项类型的名称。在一个企业中这个名称是唯一的。

group	String	
工作项类型的分组。

允许值: requirement, task, bug, issue, plan

请求示例：
{
    "name": "功能缺陷",
    "group": "bug"
}
响应示例：
{
    "id": "630da48bc9443b1aa94ce3fc",
    "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3fc",
    "name": "功能缺陷",
    "group": "bug",
    "is_system": false
}
部分更新一个工作项类型
PATCH
https://rest_api_root/v1/project/work_item_types/{work_item_type_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_type_id	String	
工作项类型的id。

请求参数
字段	类型	描述
name	String	
工作项类型的名称。在一个企业中这个名称是唯一的。

请求示例：
{
    "name": "非功能性需求"
}
响应示例：
{
    "id": "630da48bc9443b1aa94ce3df",
    "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3df",
    "name": "非功能性需求",
    "group": "requirement",
    "is_system": false
}
获取全部工作项类型列表
GET
https://rest_api_root/v1/project/work_item_types
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 10,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task",
            "is_system": true
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug",
            "is_system": true
        },
        {
            "id": "issue",
            "url": "http://rest_api_root/v1/project/work_item_types/issue",
            "name": "事务",
            "group": "issue",
            "is_system": true
        },
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
            "name": "需求",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue",
            "is_system": false
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ee",
            "name": "阶段",
            "group": "plan",
            "is_system": true
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ef",
            "name": "里程碑",
            "group": "plan",
            "is_system": true
        }
    ]
}
删除一个工作项类型
DELETE
https://rest_api_root/v1/project/work_item_types/{work_item_type_id}
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
work_item_type_id	String	
工作项类型的id。

响应示例：
{
    "id": "630da48bc9443b1aa94ce3df",
    "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3df",
    "name": "非功能性需求",
    "group": "requirement",
    "is_system": false
}
获取工作项类型方案列表
GET
https://rest_api_root/v1/project/work_item_type_plans
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id可选	String	
项目的id。查询开启本地配置的工作项类型方案时传入。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47000",
            "url": "https://rest_api_root/v1/project/work_item_type_plans/5eb623f6a70571487ea47000",
            "project_type": "scrum",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向工作项类型方案中添加一个工作项类型
POST
https://rest_api_root/v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_type_plan_id	String	
工作项类型方案的id。

请求参数
字段	类型	描述
work_item_type_id	String	
工作项类型的id。

请求示例：
{
    "work_item_type_id": "5c9b35de90ad7153c2062f18"
}
响应示例：
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
    "type_plan": {
        "id": "6abb92f18ad725395df86c45",
        "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
        "project_type": "waterfall"
    },
    "type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    },
    "sub_types": [
        {
            "id": "bug",
            "url": "https://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        }
    ]
}
部分更新工作项类型方案中的工作项类型
PATCH
https://rest_api_root/v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types/{work_item_type_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_type_plan_id	String	
工作项类型方案的id。

work_item_type_id	String	
工作项类型的id。

请求参数
字段	类型	描述
sub_type_ids	String[]	
子工作项类型id的列表，使用','分割，最多只能20个。

请求示例：
{
    "sub_type_ids": [
        "bug",
        "6385c650fef18f2d7222d15d"
    ]
}
响应示例：
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
    "type_plan": {
        "id": "6abb92f18ad725395df86c45",
        "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
        "project_type": "waterfall"
    },
    "type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    },
    "sub_types": [
        {
            "id": "bug",
            "url": "https://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        }
    ]
}
获取工作项类型方案中的工作项类型列表
GET
https://rest_api_root/v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_type_plan_id	String	
工作项类型方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
            "type_plan": {
                "id": "6abb92f18ad725395df86c45",
                "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
                "project_type": "waterfall"
            },
            "type": {
                "id": "630da48bc9443b1aa94ce3ea",
                "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
                "name": "需求",
                "group": "requirement"
            },
            "sub_types": [
                {
                    "id": "bug",
                    "url": "https://rest_api_root/v1/project/work_item_types/bug",
                    "name": "缺陷",
                    "group": "bug"
                },
                {
                    "id": "6385c650fef18f2d7222d15d",
                    "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
                    "name": "自定义",
                    "group": "issue"
                }
            ]
        }
    ]
}
在工作项类型方案中移除一个工作项类型
DELETE
https://rest_api_root/v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types/{work_item_type_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
work_item_type_plan_id	String	
工作项类型方案的id。

work_item_type_id	String	
工作项类型的id。

响应示例：
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
    "type_plan": {
        "id": "6abb92f18ad725395df86c45",
        "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
        "project_type": "waterfall"
    },
    "type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    },
    "sub_types": [
        {
            "id": "bug",
            "url": "https://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        }
    ]
}
创建一个工作项状态
POST
https://rest_api_root/v1/project/work_item_states
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
工作项状态的名称。在一个企业中这个名称是唯一的。

type	String	
工作项状态的类型。

允许值: pending, in_progress, completed, closed

请求示例：
{
    "name": "新建",
    "type": "pending"
}
响应示例：
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "pending",
    "color": "#ff7575",
    "is_system": false
}
部分更新一个工作项状态
只有非系统类型的工作项状态才能更新。

PATCH
https://rest_api_root/v1/project/work_item_states/{state_id}
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name可选	String	
工作项状态的名称。在一个企业中这个名称是唯一的。

type可选	String	
工作项状态的类型。

允许值: pending, in_progress, completed, closed

请求示例：
{
    "name": "新建",
    "type": "pending"
}
响应示例：
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "pending",
    "color": "#ff7575",
    "is_system": false
}
获取全部工作项状态列表
GET
https://rest_api_root/v1/project/work_item_states
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b35de90ad7153c2062f18",
            "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
            "name": "新建",
            "type": "pending",
            "color": "#ff7575",
            "is_system": true
        }
    ]
}
获取工作项状态方案列表
GET
https://rest_api_root/v1/project/work_item_state_plans
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id可选	String	
项目的id。查询开启本地配置的工作项状态方案时传入。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b62f18ad715335de90c20",
            "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
            "project_type": "scrum",
            "work_item_type": "story",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向状态方案中添加一个工作项状态
POST
https://rest_api_root/v1/project/work_item_state_plans/{state_plan_id}/work_item_states
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工作项状态方案的id。

请求参数
字段	类型	描述
state_id	String	
工作项状态的id。

请求示例：
{
    "state_id": "5c9b35de90ad7153c2062f18"
}
响应示例：
{
    "id": "5cc2062f189b35de90ad7153",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_states/5c9b35de90ad7153c2062f18",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    }
}
获取状态方案中的工作项状态列表
GET
https://rest_api_root/v1/project/work_item_state_plans/{state_plan_id}/work_item_states
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工作项状态方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cc2062f189b35de90ad7153",
            "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_states/5c9b35de90ad7153c2062f18",
            "state_plan": {
                "id": "5c9b62f18ad715335de90c20",
                "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            }
        }
    ]
}
在状态方案中移除一个工作项状态
DELETE
https://rest_api_root/v1/project/work_item_state_plans/{state_plan_id}/work_item_states/{state_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工作项状态方案的id。

state_id	String	
工作项状态的id。

响应示例：
{
    "id": "5cc2062f189b35de90ad7153",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_states/5c9b35de90ad7153c2062f18",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    }
}
向状态方案中添加一个工作项状态流转
POST
https://rest_api_root/v1/project/work_item_state_plans/{state_plan_id}/work_item_state_flows
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工作项状态方案的id。

请求参数
字段	类型	描述
from_state_id	String	
起始工作项状态的id。

to_state_id	String	
可更改为的工作项状态的id。

请求示例：
{
    "from_state_id": "5c9b35de90ad7153c2062f18",
    "to_state_id": "5ef85b1e9481936604da7f4c"
}
响应示例：
{
    "id": "5ef85b1e9481936604da7fcd",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_state_flows/5ef85b1e9481936604da7fcd",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "from_state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#56ABFB"
    },
    "to_state": {
        "id": "5ef85b1e9481936604da7f4c",
        "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
        "name": "进行中",
        "type": "in_progress",
        "color": "#F6C659"
    }
}
获取状态方案中的工作项状态流转列表
GET
https://rest_api_root/v1/project/work_item_state_plans/{state_plan_id}/work_item_state_flows
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工作项状态方案的id。

查询参数
字段	类型	描述
from_state_id可选	String	
上一个工作项状态的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ef85b1e9481936604da7fcd",
            "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_state_flows/5ef85b1e9481936604da7fcd",
            "state_plan": {
                "id": "5c9b62f18ad715335de90c20",
                "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "from_state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#56ABFB"
            },
            "to_state": {
                "id": "5ef85b1e9481936604da7f4c",
                "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
                "name": "进行中",
                "type": "in_progress",
                "color": "#F6C659"
            }
        }
    ]
}
在状态方案中移除一个工作项状态流转
DELETE
https://rest_api_root/v1/project/work_item_state_plans/{state_plan_id}/work_item_state_flows/{flow_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
state_plan_id	String	
工作项状态方案的id。

flow_id	String	
工作项状态流转的id。

响应示例：
{
    "id": "5ef85b1e9481936604da7fcd",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_state_flows/5ef85b1e9481936604da7fcd",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "from_state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#56ABFB"
    },
    "to_state": {
        "id": "5ef85b1e9481936604da7f4c",
        "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
        "name": "进行中",
        "type": "in_progress",
        "color": "#F6C659"
    }
}
创建一个工作项属性
POST
https://rest_api_root/v1/project/work_item_properties
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
工作项属性的名称。在一个企业中这个名称是唯一的。

type	String	
工作项属性的类型。

允许值: text, textarea, select, multi_select, cascade_select, cascade_multi_select, member, members, date, number, progress, rate, link

options可选	Object[]	
选择项列表。当工作项属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。

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
            "_id": "5efb1859110533727a82c604",
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
    "url": "https://rest_api_root/v1/project/work_item_properties/severity",
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
    "url": "https://rest_api_root/v1/project/work_item_properties/jiliandanxuan",
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
部分更新一个工作项属性
PATCH
https://rest_api_root/v1/project/work_item_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_id	String	
工作项属性的id。

请求参数
字段	类型	描述
name可选	String	
工作项属性的名称。在一个企业中这个名称是唯一的。

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
        },
        {
            "text": "子-create",
            "parent_id": "64b9f741eec7fd94e63b411f"
        }
    ]
}
响应示例（select）：
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/project/work_item_properties/severity",
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
    "url": "https://rest_api_root/v1/project/work_item_properties/jiliandanxuan",
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
        },
        {
            "_id": "64b9f741eec7fd94e63b411g",
            "text": "子-create",
            "parent_id": "64b9f741eec7fd94e63b411f"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
获取全部工作项属性列表
GET
https://rest_api_root/v1/project/work_item_properties
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "severity",
            "url": "https://rest_api_root/v1/project/work_item_properties/severity",
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
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/project/work_item_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false,
            "select_all_level": false,
            "display_all_level": false,
            "display_separator": null
        }
    ]
}
获取工作项属性方案列表
GET
https://rest_api_root/v1/project/work_item_property_plans
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
project_id可选	String	
项目的id。查询开启本地配置的工作项属性方案时传入。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5f8a21f18ef715265de90c21",
            "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
            "project_type": "scrum",
            "work_item_type": "story",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
向属性方案中添加一个工作项属性
POST
https://rest_api_root/v1/project/work_item_property_plans/{property_plan_id}/work_item_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
工作项属性方案的id。

请求参数
字段	类型	描述
property_id	String	
工作项属性的id。

请求示例：
{
    "property_id": "severity"
}
响应示例：
{
    "id": "severity",
    "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/severity",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "property": {
        "id": "severity",
        "url": "https://rest_api_root/v1/project/work_item_properties/severity",
        "name": "严重程度",
        "type": "select"
    }
}
获取属性方案中的工作项属性列表
GET
https://rest_api_root/v1/project/work_item_property_plans/{property_plan_id}/work_item_properties
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
工作项属性方案的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "severity",
            "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/severity",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "property": {
                "id": "severity",
                "url": "https://rest_api_root/v1/project/work_item_properties/severity",
                "name": "严重程度",
                "type": "select"
            }
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/identifier",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "property": {
                "id": "identifier",
                "url": "https://rest_api_root/v1/project/work_item_properties/identifier",
                "name": "编号",
                "type": "text"
            }
        }
    ]
}
在属性方案中移除一个工作项属性
DELETE
https://rest_api_root/v1/project/work_item_property_plans/{property_plan_id}/work_item_properties/{property_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
property_plan_id	String	
工作项属性方案的id。

property_id	String	
工作项属性的id。

响应示例：
{
    "id": "severity",
    "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/severity",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "property": {
        "id": "severity",
        "url": "https://rest_api_root/v1/project/work_item_properties/severity",
        "name": "严重程度",
        "type": "select"
    }
}
开启工作项本地配置
POST
https://rest_api_root/v1/project/work_item_plans
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
project_id	String	
项目的id。

请求示例：
{
    "project_id": "5eb623f6a70571487ea47000"
}
响应示例：
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/work_item_plans/5eb623f6a70571487ea47000",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    }
}