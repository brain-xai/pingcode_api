全局
个人
用于查看和管理个人的基本信息。

获取个人信息
GET
https://rest_api_root/v1/myself
权限: 用户令牌

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "email": "terry@pingcode.com",
    "mobile": "15000000000",
    "status": "enabled",
    "preferences": {
        "locale": "zh-cn",
        "timezone": "Asia/Shanghai"
    },
    "permissions": {
        "agile_create_project": true,
        "agile_configuration": true,
        "create_user": true
    }
}
组织
企业
用于查看企业的基本信息。

资源属性
字段	类型	描述
id	String	
企业的id。

url	String	
企业的资源地址。

name	String	
企业的名称。

secondary_domain	String	
企业的二级域名。

全量数据示例：
{
    "id": "56ba35de87ad7153c2062f65",
    "url": "https://rest_api_root/v1/directory/team",
    "name": "YCtech",
    "secondary_domain": "yctech"
}
引用数据示例：
{
    "id": "56ba35de87ad7153c2062f65",
    "url": "https://rest_api_root/v1/directory/team",
    "name": "YCtech",
    "secondary_domain": "yctech"
}
获取企业信息
GET
https://rest_api_root/v1/directory/team
权限: 企业令牌/用户令牌

响应示例：
{
    "id": "56ba35de87ad7153c2062f65",
    "url": "https://rest_api_root/v1/directory/team",
    "name": "YCtech",
    "secondary_domain": "yctech"
}
企业成员
用于查看和管理企业成员的基本信息。
资源地址：{GET} https://rest_api_root/v1/directory/users/{user_id}

资源属性
字段	类型	描述
id	String	
企业成员的id。

url	String	
企业成员的资源地址。

name	String	
企业成员的名称。

display_name	String	
企业成员的显示名。

avatar	String	
企业成员的头像地址。

department	Object	
企业成员的部门。

job	Object	
企业成员的职位。

email	String	
企业成员的邮箱地址。

mobile	String	
企业成员的手机号。

status	String	
企业成员的状态。

允许值: enabled, disabled

employee_number	String	
企业成员的工号。

全量数据示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "department": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
        "name": "技术支持"
    },
    "job": {
        "id": "6440c881c56f557eb1aff6e5",
        "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
        "name": "后端工程师"
    },
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "enabled",
    "employee_number": "zxv"
}
引用数据示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
}
创建一个企业成员
POST
https://rest_api_root/v1/directory/users
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
企业成员的名称，在一个企业中这个名称是唯一的。

display_name	String	
企业成员的显示名。

email可选	String	
企业成员的邮箱地址，在一个企业中这个邮箱地址是唯一的，邮箱地址和手机号其中一个必填。

mobile可选	String	
企业成员的手机号，在一个企业中这个手机号是唯一，邮箱地址和手机号其中一个必填。

password可选	String	
企业成员的密码，长度为6～200的字符串(包含6和200)。

department_id可选	String	
企业成员的部门id。

job_id可选	String	
企业成员的职位id。

employee_number可选	String	
企业成员的工号。

请求示例：
{
    "name": "john",
    "display_name": "John",
    "email": "john@email.com",
    "mobile": "15000000000",
    "password": "123456",
    "department_id": "6422711c3f12e6c1e46d40e6",
    "job_id": "6440c881c56f557eb1aff6e5",
    "employee_number": "zxv"
}
响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "department": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
        "name": "技术支持"
    },
    "job": {
        "id": "6440c881c56f557eb1aff6e5",
        "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
        "name": "后端工程师"
    },
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "init",
    "employee_number": "zxv"
}
部分更新一个企业成员
PATCH
https://rest_api_root/v1/directory/users/{user_id}
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name可选	String	
企业成员的名称，在一个企业中这个名称是唯一的。

display_name可选	String	
企业成员的显示名。

email可选	String	
企业成员的邮箱地址，在一个企业中这个邮箱地址是唯一的。

mobile可选	String	
企业成员的手机号，在一个企业中这个手机号是唯一的。

status可选	String	
企业成员的状态。

允许值: enabled, disabled

employee_number可选	String	
企业成员的工号。

department_id可选	String	
企业成员的部门id。

job_id可选	String	
企业成员的职位id。

请求示例：
{
    "name": "john",
    "display_name": "John",
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "enabled",
    "employee_number": "zxv",
    "department_id": "6422711c3f12e6c1e46d40e6",
    "job_id": "6440c881c56f557eb1aff6e5"
}
响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "department": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
        "name": "技术支持"
    },
    "job": {
        "id": "6440c881c56f557eb1aff6e5",
        "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
        "name": "后端工程师"
    },
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "enabled",
    "employee_number": "zxv"
}
批量更新企业成员属性
PATCH
https://rest_api_root/v1/directory/users/bulk
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
user_ids	String[]	
企业成员的id数组，不能包含自己和团队拥有者。

property_name	String	
需要更新的企业成员属性名，目前仅支持：status（可选值为：enabled、disabled）

property_value	String	
需要更新的企业成员属性值。

请求示例：
{
    "user_ids": [
        "a0417f68e846aae315c85d24643678a9",
        "a0417f68e846aae315c85d24643678a8"
    ],
    "property_name": "status",
    "property_value": "enabled"
}
响应示例：
[
    {
        "state": "success",
        "user_id": "a0417f68e846aae315c85d24643678a9"
    },
    {
        "state": "failure",
        "user_id": "a0417f68e846aae315c85d24643678a8",
        "message": "failure reason.."
    }
]
获取企业成员列表
GET
https://rest_api_root/v1/directory/users
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
name可选	String	
企业成员的名称，在一个企业中这个名称是唯一的。

keywords可选	String	
关键词模糊搜索，支持姓名、用户名。

department_ids可选	String	
企业成员的部门id，使用','分割，最多只能20个。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
            "name": "john",
            "display_name": "John",
            "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
            "department": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
                "name": "技术支持"
            },
            "job": {
                "id": "6440c881c56f557eb1aff6e5",
                "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
                "name": "后端工程师"
            },
            "email": "john@email.com",
            "mobile": "15000000000",
            "status": "enabled",
            "employee_number": "zxv"
        }
    ]
}
部门
用于查看和管理企业的部门信息。
资源地址：{GET} https://rest_api_root/v1/directory/departments/{department_id}

资源属性
字段	类型	描述
id	String	
部门的id。

url	String	
部门的资源地址。

name	String	
部门的名称。

head	Object	
部门的负责人。

parent	Object	
父部门。

全量数据示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
引用数据示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持"
}
创建一个部门
POST
https://rest_api_root/v1/directory/departments
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
部门的名称，在一个企业中这个名称是唯一的。

parent_id可选	String	
父部门的id。

head_id可选	String	
部门负责人的id。

请求示例：
{
    "name": "技术支持",
    "parent_id": "6422711c3f12e6c1e46d40e2",
    "head_id": "70e9933e5e7948779b9b8978b6489038"
}
响应示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
部分更新一个部门
PATCH
https://rest_api_root/v1/directory/departments/{department_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
department_id	String	
部门id。

请求参数
字段	类型	描述
name可选	String	
部门的名称，在一个企业中这个名称是唯一的。

parent_id可选	String	
父部门的id。

head_id可选	String	
部门负责人的id。

请求示例：
{
    "name": "技术支持",
    "parent_id": "6422711c3f12e6c1e46d40e2",
    "head_id": "70e9933e5e7948779b9b8978b6489038"
}
响应示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
获取部门列表
GET
https://rest_api_root/v1/directory/departments
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e6",
            "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
            "name": "技术支持",
            "head": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "parent": {
                "id": "6422711c3f12e6c1e46d40e2",
                "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
                "name": "技术中心"
            }
        }
    ]
}
删除一个部门
DELETE
https://rest_api_root/v1/directory/departments/{department_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
department_id	String	
部门id。

响应示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
团队
用于查看企业的团队信息。
资源地址：{GET} https://rest_api_root/v1/directory/groups/{group_id}

资源属性
字段	类型	描述
id	String	
团队的id。

url	String	
团队的资源地址。

name	String	
团队的名称。

visibility	String	
团队的可见性。

允许值: private, public

description	String	
团队的描述。

全量数据示例：
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team",
    "visibility": "private",
    "description": "This is Open Team."
}
引用数据示例：
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team"
}
创建一个团队
POST
https://rest_api_root/v1/directory/groups
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
name	String	
团队的名称，在一个企业中这个名称是唯一的。

visibility可选	String	
团队的可见范围。

默认值: private

允许值: private, public

description可选	String	
团队的描述。

请求示例：
{
    "name": "Open Team",
    "visibility": "private",
    "description": "This is Open Team."
}
响应示例：
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team",
    "visibility": "private",
    "description": "This is Open Team."
}
部分更新一个团队
PATCH
https://rest_api_root/v1/directory/groups/{group_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
group_id	String	
团队id。

请求参数
字段	类型	描述
name可选	String	
团队的名称，在一个企业中这个名称是唯一的。

visibility可选	String	
团队的可见范围。

允许值: private, public

description可选	String	
团队的描述。

请求示例：
{
    "name": "Open Team Update",
    "visibility": "public",
    "description": "This is Update Open Team."
}
响应示例：
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team Update",
    "visibility": "public",
    "description": "This is Update Open Team."
}
获取团队列表
GET
https://rest_api_root/v1/directory/groups
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
            "name": "Open Team",
            "visibility": "private",
            "description": "This is Open Team."
        },
        {
            "id": "64ca0f67cb78a0a80e1a999e",
            "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
            "name": "PingCode",
            "visibility": "public",
            "description": "This is PingCode."
        }
    ]
}
向团队中添加一个成员
POST
https://rest_api_root/v1/directory/groups/{group_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
group_id	String	
团队id。

请求参数
字段	类型	描述
user_id	String	
用户id。

role	String	
团队角色。

允许值: manager, member

请求示例：
{
    "user_id": "a0417f68e846aae315c85d24643678a9",
    "role": "manager"
}
响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e/members/a0417f68e846aae315c85d24643678a9",
    "group": {
        "id": "64ca0f67cb78a0a80e1a999e",
        "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
        "name": "PingCode"
    },
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": "manager"
}
获取团队中的成员列表
GET
https://rest_api_root/v1/directory/groups/{group_id}/members
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
group_id	String	
团队id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e/members/a0417f68e846aae315c85d24643678a9",
            "group": {
                "id": "64ca0f67cb78a0a80e1a999e",
                "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
                "name": "PingCode"
            },
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": "manager"
        }
    ]
}
在团队中移除一个成员
DELETE
https://rest_api_root/v1/directory/groups/{group_id}/members/{member_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
group_id	String	
团队id。

member_id	String	
团队成员id。

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e/members/a0417f68e846aae315c85d24643678a9",
    "group": {
        "id": "64ca0f67cb78a0a80e1a999e",
        "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
        "name": "PingCode"
    },
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": "manager"
}
角色
用于查看企业的角色信息。
资源地址：{GET} https://rest_api_root/v1/directory/roles/{role_id}

资源属性
字段	类型	描述
id	String	
角色的id。

url	String	
角色的资源地址。

name	String	
角色的名称。

is_system	Boolean	
角色是否为系统内置。

允许值: true, false

全量数据示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
    "name": "管理员",
    "is_system": true
}
引用数据示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
    "name": "管理员"
}
获取角色列表
GET
https://rest_api_root/v1/directory/roles
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e6",
            "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
            "name": "管理员",
            "is_system": true
        }
    ]
}
职位
用于查看企业的职位信息。
资源地址：{GET} https://rest_api_root/v1/directory/jobs/{job_id}

资源属性
字段	类型	描述
id	String	
职位的id。

url	String	
职位的资源地址。

name	String	
职位的名称。

is_system	Boolean	
职位是否为系统内置。

允许值: true, false

全量数据示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/jobs/6422711c3f12e6c1e46d40e6",
    "name": "技术总监",
    "is_system": true
}
引用数据示例：
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/jobs/6422711c3f12e6c1e46d40e6",
    "name": "技术总监"
}
获取职位列表
GET
https://rest_api_root/v1/directory/jobs
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e6",
            "url": "https://rest_api_root/v1/directory/jobs/6422711c3f12e6c1e46d40e6",
            "name": "技术总监",
            "is_system": true
        }
    ]
}
安全
日志
用于查看企业的日志信息。

获取登录日志列表
GET
https://rest_api_root/v1/security/login_logs
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
logged_between	String	
登录时间介于的时间范围，通过','分割起始时间。

user_ids可选	String	
成员id的列表，使用','分割，最多只能20个。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5fca7b74efab845a2fa53181",
            "url": "https://rest_api_root/v1/security/login_logs/5fca7b74efab845a2fa53181",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "login_method": "账号密码",
            "region": "北京海淀区",
            "ip": "45.251.20.42",
            "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
            "created_at": 1676454024
        }
    ]
}
获取审计日志列表
GET
https://rest_api_root/v1/security/audit_logs
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
operated_between	String	
操作时间介于的时间范围，通过','分割起始时间。

operated_bys可选	String	
操作人的列表，使用','分割，最多只能20个。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5fca7b74efab845a2fa53181",
            "url": "https://rest_api_root/v1/security/audit_logs/5fca7b74efab845a2fa53181",
            "operated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "operated_at": 1676454024,
            "operate_object": "规则",
            "application": "自动化",
            "ip": "45.251.20.42",
            "summary": "修改规则",
            "detail": "规则：规则1\n类型：自动化规则\n描述：-"
        }
    ]
}
通用
评论
用于查看和管理评论的基本信息。
资源地址：{GET} https://rest_api_root/v1/comments/{comment_id}

资源属性
字段	类型	描述
id	String	
评论的id。

url	String	
评论的资源地址。

content	String	
评论的内容。被删除的评论content字段会被置空，is_deleted字段为1。

attachment_count	Number	
评论的附件数量。

attachments	Object[]	
评论的附件列表。

created_at	Number	
评论的创建时间。

created_by	Object	
评论的创建人。

is_deleted	Number	
评论是否被删除。

允许值: 0, 1

全量数据示例：
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "attachment_count": 1,
    "attachments": [
        {
            "id": "5da588ca84c7377a5d327e6d",
            "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
            "title": "这是一个代码片段",
            "size": 16,
            "type": "snippet"
        }
    ],
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 0
}
引用数据示例：
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "attachment_count": 1,
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 0
}
创建一个评论
POST
https://rest_api_root/v1/comments
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
principal_type	String	
评论主体的类型。

允许值: work_item, test_run, test_case, ticket, idea, page

principal_id可选	String	
评论主体的id。

review_id可选	String	
评论评审的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。

content	String	
评论的内容。

created_at可选	Number	
评论的创建时间。

created_by可选	String	
评论的创建人。

请求示例：
{
    "principal_type": "work_item",
    "principal_id": "5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "created_at": 1565255712,
    "created_by": "a0417f68e846aae315c85d24643678a9"
}
响应示例：
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "attachment_count": 0,
    "attachments": [],
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 0
}
获取评论列表
GET
https://rest_api_root/v1/comments?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
principal_type	String	
评论主体的类型。

允许值: work_item, test_run, test_case, ticket, idea, page

principal_id可选	String	
评论主体的id。

review_id可选	String	
评论评审的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "59f72dfaeadb5b5197b7da6d",
            "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
            "content": "这是一个工作项评论",
            "attachment_count": 0,
            "attachments": [],
            "created_at": 1565255712,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_deleted": 0
        }
    ]
}
删除一个评论
DELETE
https://rest_api_root/v1/comments/{comment_id}?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
comment_id	String	
评论的id。

查询参数
字段	类型	描述
principal_type	String	
评论主体的类型。

允许值: work_item, test_run, test_case, ticket, idea, page

principal_id可选	String	
评论主体的id。

review_id可选	String	
评论评审的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。

响应示例：
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "",
    "attachment_count": 0,
    "attachments": [],
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 1
}
附件
用于查看和管理附件的基本信息。
资源地址：{GET} https://rest_api_root/v1/attachments/{attachment_id}

资源属性
字段	类型	描述
id	String	
附件的id。

url	String	
附件的资源地址。

title	String	
附件的标题。

size	Number	
附件的大小。

type	String	
附件的类型。

允许值: file, snippet

file_type可选	String	
文件的类型。当附件的类型type为file时，该字段会返回。

允许值: image, other

ext可选	String	
文件的拓展名。当附件的类型type为file时，该字段会返回。

download_url可选	String	
文件的下载地址。当附件的类型type为file时，该字段会返回。

format可选	String	
代码的格式。当附件的类型type为snippet时，该字段会返回。

允许值: clike, css, dart, django, dockerfile, go, markdown, nginx, python, php, shell, sql, swift, html, javascript, jsx, pascal, sass, stylus, vue, yaml, haskell

content可选	String	
代码的内容。当附件的类型type为snippet时，该字段会返回。

line可选	String	
代码的行数。当附件的类型type为snippet时，该字段会返回。

created_at	Number	
附件的创建时间。

created_by	Object	
附件的创建人。

全量数据示例（文件）:
{
    "id": "5da588ca84c7377a5d327e6c",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6c?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "title": "这是一个图片类型的附件",
    "size": 1024,
    "type": "file",
    "file_type": "image",
    "ext": "png",
    "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a.png",
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
全量数据示例（代码段）:
{
    "id": "5da588ca84c7377a5d327e6d",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
    "title": "这是一个代码片段",
    "size": 16,
    "type": "snippet",
    "format": "javascript",
    "content": "const a = 'abc';",
    "line": 1,
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
引用数据示例：
{
    "id": "5da588ca84c7377a5d327e6d",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
    "title": "这是一个代码片段",
    "size": 16,
    "type": "snippet"
}
上传一个文件
POST
https://rest_api_root/v1/attachments?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

Header
字段	类型	描述
content-type	String	
multipart/form-data。

查询参数
字段	类型	描述
principal_type	String	
附件主体的类型。

允许值: work_item, work_item_deliverable, test_case, test_run, idea, ticket, page

principal_id	String	
附件主体的id。

comment_id可选	String	
附件主体的评论id。当需要向附件主体的评论上传文件或者代码段时，需要传入该参数。

请求参数 form-data
字段	类型	描述
title	String	
这是一个图片类型的附件.png

file	File	
/Users/ping-code/这是一个图片类型的附件.png

响应示例：
{
    "id": "5da588ca84c7377a5d327e6c",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6c?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "title": "这是一个图片类型的附件",
    "size": 1024,
    "type": "file",
    "file_type": "image",
    "ext": "png",
    "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a.png",
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
上传一个代码段
POST
https://rest_api_root/v1/attachments
权限: 企业令牌/用户令牌

Header
字段	类型	描述
content-type	String	
application/json。

请求参数
字段	类型	描述
principal_type	String	
附件主体的类型。

允许值: work_item, test_case, test_run, idea, ticket, page

principal_id	String	
附件主体的id。

comment_id可选	String	
附件主体的评论id。当需要向附件主体的评论上传文件或者代码段时，需要传入该参数。

title	String	
代码段的标题。

format	String	
代码段的语言。

允许值: clike, css, dart, django, dockerfile, go, markdown, nginx, python, php, shell, sql, swift, html, javascript, jsx, pascal, sass, stylus, vue, yaml, haskell

content	String	
代码段的内容。

请求示例：
{
    "principal_type": "work_item",
    "principal_id": "5edca524cad2fa1125cb0630",
    "comment_id": "59f72dfaeadb5b5197b7da6d",
    "title": "这是一个代码片段",
    "format": "javascript",
    "content": "const a = 'abc';"
}
响应示例：
{
    "id": "5da588ca84c7377a5d327e6d",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
    "title": "这是一个代码片段",
    "size": 16,
    "type": "snippet",
    "format": "javascript",
    "content": "const a = 'abc';",
    "line": 1,
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
获取附件列表
GET
https://rest_api_root/v1/attachments?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
principal_type	String	
附件主体的类型。

允许值: work_item, test_case, test_run, idea, ticket, page

principal_id	String	
附件主体的id。

comment_id可选	String	
附件主体的评论id。当需要获取附件主体的评论附件时，需要传入该参数。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5da588ca84c7377a5d327e6d",
            "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
            "title": "这是一个代码片段",
            "size": 16,
            "type": "snippet",
            "format": "javascript",
            "content": "const a = 'abc';",
            "line": 1,
            "created_at": 1583290347,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "5da588ca84c7377a5d327e6f",
            "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6f?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
            "title": "这是一个图片类型的附件",
            "size": 1024,
            "type": "file",
            "file_type": "image",
            "ext": "png",
            "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839b.png",
            "created_at": 1583290347,
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
删除一个附件
DELETE
https://rest_api_root/v1/attachments/{attachment_id}?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
attachment_id	String	
附件的id。

查询参数
字段	类型	描述
principal_type	String	
附件主体的类型。

允许值: work_item, test_case, test_run, idea, ticket, page

principal_id	String	
附件主体的id。

comment_id可选	String	
附件主体的评论id。当需要获取附件主体的评论附件时，需要传入该参数。

响应示例：
{
    "id": "5da588ca84c7377a5d327e6c",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6c?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "title": "这是一个图片类型的附件",
    "size": 1024,
    "type": "file",
    "file_type": "image",
    "ext": "png",
    "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a.png",
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
关注人
用于查看和管理关注人的基本信息。
资源地址：{GET} https://rest_api_root/v1/participants/{participant_id}

资源属性
字段	类型	描述
id	String	
关注人的id。

url	String	
关注人的资源地址。

type	String	
关注人的类型。

允许值: user, user_group

user可选	Object	
关注的用户。当type为user时，该字段返回。

user_group可选	Object	
关注的团队。当type为user_group时，该字段返回。

全量数据示例（用户类型）:
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
全量数据示例（评审用户类型）:
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
全量数据示例（团队类型）:
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user_group",
    "user_group": {
        "id": "63c8fb32729dee3334d96af7",
        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
        "name": "Open Team"
    }
}
引用数据示例（用户类型）:
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
引用数据示例（团队类型）:
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user_group",
    "user_group": {
        "id": "63c8fb32729dee3334d96af7",
        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
        "name": "Open Team"
    }
}
添加一个关注人
POST
https://rest_api_root/v1/participants
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
principal_type	String	
关注人主体的类型。

允许值: work_item, test_case, idea, ticket, page

principal_id可选	String	
关注人主体的id。

review_id可选	String	
关注人评审主体的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。

type	String	
关注人的类型。

允许值: user, user_group

participant_id	String	
关注人的id。用户的id或者团队的id。

请求示例（工作项）：
{
    "principal_type": "work_item",
    "principal_id": "63e1bf51760505c8795ebccc",
    "type": "user",
    "participant_id": "a0417f68e846aae315c85d24643678a9"
}
请求示例（产品需求评审）：
{
    "principal_type": "idea",
    "review_id": "6f168f764eba01a5278b87cd",
    "type": "user",
    "participant_id": "a0417f68e846aae315c85d24643678a9"
}
响应示例（工作项）：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
响应示例（产品需求评审）：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
获取关注人列表
GET
https://rest_api_root/v1/participants?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
principal_type	String	
关注人主体的类型。

允许值: work_item, test_case, idea, ticket, page

principal_id可选	String	
关注人主体的id。

review_id可选	String	
关注人评审主体的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
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
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ]
}
移除一个关注人
DELETE
https://rest_api_root/v1/participants/{participant_id}?principal_type={principal_type}&principal_id={principal_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
participant_id	String	
关注人的id。用户的id或者团队的id。

查询参数
字段	类型	描述
principal_type	String	
关注人主体的类型。

允许值: work_item, test_case, idea, ticket, page

principal_id可选	String	
关注人主体的id。

review_id可选	String	
注人评审主体的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。

响应示例：
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
关联
用于查看和管理关联的基本信息。
资源地址：{GET} https://rest_api_root/v1/relations/{relation_id}

资源属性
字段	类型	描述
id	String	
关联的id。

url	String	
关联的资源地址。

principal_type	String	
关联主体的类型。

允许值: idea, ticket, work_item, test_plan, test_run, test_case, page

principal	Object	
关联的主体。

target_type	String	
关联的目标类型。

允许值: ticket, work_item, test_case, idea, page

target	Object	
关联的目标。

全量数据示例（工作项关联需求）：
{
    "id": "653b1b4a3113be5bb040e4c5",
    "url": "https://rest_api_root/v1/relations/653b1b4a3113be5bb040e4c5",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_type": "idea",
    "target": {
        "id": "64b4d70ba368e6594360ea24",
        "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
        "identifier": "SLC-1",
        "title": "示例需求",
        "short_id": "Ogf1EYey",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
    }
}
全量数据示例（测试计划关联工作项）：
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
    "principal_type": "test_plan",
    "principal": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870
    },
    "target_type": "work_item",
    "target": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
全量数据示例（工单关联测试用例）：
{
    "id": "6552d7ab1bffb252b82645ba",
    "url": "https://rest_api_root/v1/relations/6552d7ab1bffb252b82645ba",
    "principal_type": "ticket",
    "principal": {
        "id": "63eca888a0a13a3efc8d4a43",
        "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
        "identifier": "SLC-T1",
        "title": "希望新增支持第三方账号注册",
        "short_id": "pdMUgQzZ",
        "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
    },
    "target_type": "test_case",
    "target": {
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
}
全量数据示例（工作项关联页面）：
{
    "id": "6552d9da1bffb252b8276cf7",
    "url": "https://rest_api_root/v1/relations/6552d9da1bffb252b8276cf7",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_type": "page",
    "target": {
        "id": "63e1bf51760505c8795ebccc",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
        "name": "示例页面",
        "type": "document",
        "short_id": "5-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
    }
}
引用数据示例:
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2"
}
创建一个关联
POST
https://rest_api_root/v1/relations
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
principal_type	String	
关联主体的类型。关联主体的类型和关联的目标类型需要搭配使用，比如：
需求关联工单，principal_type为idea，target_type为ticket；
需求关联工作项，principal_type为idea，target_type为work_item；
需求关联测试用例，principal_type为idea，target_type为test_case；
需求关联需求，principal_type为idea，target_type为idea；
需求关联页面，principal_type为idea，target_type为page；
工单关联需求，principal_type为ticket，target_type为idea；
工单关联工作项，principal_type为ticket，target_type为work_item；
工单关联工单，principal_type为ticket，target_type为ticket；
工单关联页面，principal_type为ticket，target_type为page；
工作项关联测试用例，principal_type为work_item，target_type为test_case；
工作项关联需求，principal_type为work_item，target_type为idea；
工作项关联工单，principal_type为work_item，target_type为ticket；
工作项关联页面，principal_type为work_item，target_type为page；
测试计划关联缺陷，principal_type为test_plan，target_type为work_item；
执行用例关联缺陷，principal_type为test_run，target_type为work_item；
测试用例关联需求，principal_type为test_case，target_type为idea；
测试用例关联工作项，principal_type为test_case，target_type为work_item；
测试用例关联页面，principal_type为test_case，target_type为page；
页面关联需求，暂不开放；
页面关联工单，暂不开放；
页面关联工作项，暂不开放；
页面关联测试用例，暂不开放；

principal_id	String	
关联主体的id。

target_type	String	
关联的目标类型。

target_id	String	
关联的目标id。

请求示例：
{
    "principal_id": "547000eb6a70571487623fea",
    "principal_type": "test_run",
    "target_type": "work_item",
    "target_id": "5edca524cad2fa1125cb0630"
}
响应示例：
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
    "principal_type": "test_run",
    "principal": {
        "id": "547000eb6a70571487623fea",
        "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
        "status": "failure",
        "short_id": "Aq1u38yX",
        "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX"
    },
    "target_type": "work_item",
    "target": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
获取关联列表
GET
https://rest_api_root/v1/relations?principal_type={principal_type}&principal_id={principal_id}&target_type={target_type}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
principal_type	String	
关联主体的类型。关联主体的类型和关联的目标类型需要搭配使用，比如：
需求关联工单，principal_type为idea，target_type为ticket；
需求关联工作项，principal_type为idea，target_type为work_item；
需求关联测试用例，principal_type为idea，target_type为test_case；
需求关联需求，principal_type为idea，target_type为idea；
需求关联页面，principal_type为idea，target_type为page；
工单关联需求，principal_type为ticket，target_type为idea；
工单关联工作项，principal_type为ticket，target_type为work_item；
工单关联工单，principal_type为ticket，target_type为ticket；
工单关联页面，principal_type为ticket，target_type为page；
工作项关联测试用例，principal_type为work_item，target_type为test_case；
工作项关联需求，principal_type为work_item，target_type为idea；
工作项关联工单，principal_type为work_item，target_type为ticket；
工作项关联页面，principal_type为work_item，target_type为page；
测试计划关联缺陷，principal_type为test_plan，target_type为work_item；
执行用例关联缺陷，principal_type为test_run，target_type为work_item
测试用例关联需求，principal_type为test_case，target_type为idea；
测试用例关联工作项，principal_type为test_case，target_type为work_item
测试用例关联页面，principal_type为test_case，target_type为page；
页面关联需求，principal_type为page，target_type为idea；
页面关联工单，principal_type为page，target_type为ticket；
页面关联工作项，principal_type为page，target_type为work_item；
页面关联测试用例，principal_type为page，target_type为test_case；

principal_id	String	
关联主体的id。

target_type	String	
关联的目标类型。

Success-Response
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64b4d70ba368e6594360ea79",
            "url": "https://rest_api_root/v1/relations/64b4d70ba368e6594360ea79",
            "principal_type": "idea",
            "principal": {
                "id": "64b4d70ba368e6594360ea24",
                "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
                "identifier": "SLC-1",
                "title": "示例需求",
                "short_id": "Ogf1EYey",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
            },
            "target_type": "ticket",
            "target": {
                "id": "63eca888a0a13a3efc8d4a43",
                "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
                "identifier": "SLC-T1",
                "title": "希望新增支持第三方账号注册",
                "short_id": "pdMUgQzZ",
                "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
            }
        }
    ]
}
删除一个关联
DELETE
https://rest_api_root/v1/relations/{relation_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
relation_id	String	
关联的id。

响应示例：
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
    "principal_type": "test_plan",
    "principal": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870
    },
    "target_type": "work_item",
    "target": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
工时
用于查看和管理工时的基本信息。
资源地址：{GET} https://rest_api_root/v1/workloads/{workload_id}

资源属性
字段	类型	描述
id	String	
工时的id。

url	String	
工时的资源地址。

principal_type	String	
工时主体的类型。

允许值: work_item

principal	Object	
工时的主体。

type	Object	
工时的类型。

duration	Number	
工时的时长。单位是小时，数值可以是为0-24之间，最多包含一位小数的正数。

review_state	String	
工时的评审状态。

允许值: no_review, pending, in_progress, approved, rejected, terminated

description	String	
工时的说明。

report_at	Number	
工时的登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。

report_by	Object	
工时的登记人。

created_at	Number	
工时的创建日期。该值为十位数字组成的时间戳。

created_by	Object	
工时的创建人。

全量数据示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
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
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
引用数据示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
创建一个工时
POST
https://rest_api_root/v1/workloads
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
principal_type	String	
工时主体的类型。

允许值: work_item, idea, test_case

principal_id	String	
工时主体的id。

type_id可选	String	
工时类型的id。

duration	Number	
工时的时长。单位是小时，数值可以是为0-24之间，最多包含一位小数的正数。

report_at	Number	
工时的登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。

report_by_id可选	String	
工时的登记人，企业鉴权时必填。个人鉴权时不需要传递，即使传递了也会被忽略。

recorded_at可选	String	
工时的登记时间，默认是当前时间。

description可选	String	
工时的说明。

请求示例：
{
    "principal_id": "5edca524cad2fa1125cb0630",
    "principal_type": "work_item",
    "type_id": "5a86eaf6a72585327ea46fge0",
    "duration": 8,
    "report_at": 1593290347,
    "report_by_id": "a0417f68e846aae315c85d24643678a9",
    "recorded_at": 1593290347,
    "description": "这是一个工时"
}
响应示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
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
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
部分更新一个工时
用户令牌只能更新属于用户自己登记的工时记录。

PATCH
https://rest_api_root/v1/workloads/{workload_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
workload_id	String	
工时的id。

请求参数
字段	类型	描述
type_id可选	String	
工时类型的id。

duration可选	Number	
工时的时长。单位是小时，数值可以是为0-24之间，最多包含一位小数的正数。

report_at可选	Number	
工时的登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。

description可选	String	
工时的说明。

请求示例：
{
    "type_id": "5a86eaf6a72585327ea46fge0",
    "duration": 8,
    "report_at": 1593290347,
    "description": "这是一个工时"
}
响应示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
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
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
获取工时列表
GET
https://rest_api_root/v1/workloads
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
principal_type可选	String	
工时主体的类型。当查询参数含有pilot_id或principal_id时，principal_type参数必填。

允许值: idea, work_item, test_case

pilot_id可选	String	
工时主体所在产品、项目或测试库的id。使用该参数过滤数据时，登记日期查询的起始时间和登记日期查询的结束时间的跨度最大为3个月。

principal_id可选	String	
工时主体的id。

start_at可选	Number	
登记日期查询的起始时间。开始时间会转换为对应日期的开始时间点。开始时间和结束时间须同时存在。

end_at可选	Number	
登记日期查询的结束时间。结束时间会转换为对应日期的结束时间点。

report_by_id可选	String	
登记人的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5f168f764eba01a5278b87cd",
            "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
            "principal_type": "work_item",
            "principal": {
                "id": "5edca524cad2fa1125cb0630",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
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
            "type": {
                "id": "5a86eaf6a72585327ea46fge0",
                "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
                "name": "研发"
            },
            "duration": 8,
            "review_state": "approved",
            "description": "这是一个工时",
            "report_at": 1593290347,
            "report_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "created_at": 1593290347,
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
删除一个工时
用户令牌只能删除用户自己登记的工时记录。

DELETE
https://rest_api_root/v1/workloads/{workload_id}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
workload_id	String	
工时的id。

响应示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
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
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
获取工时类型列表
GET
https://rest_api_root/v1/workload_types
权限: 企业令牌/用户令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5a86eaf6a72585327ea46fge0",
            "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
            "name": "研发"
        }
    ]
}
评审
用于查看和管理评审的基本信息。
资源地址：{GET} https://rest_api_root/v1/reviews/{review_id}?principal_type={principal_type}

资源属性
字段	类型	描述
id	String	
评审的 id。

url	String	
评审的资源地址。

identifier	String	
评审的标识符。

title	String	
评审的标题。

status	String	
评审状态。

允许值: pending, in_progress, completed, repealed

principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

short_id	String	
评审的短id。

html_url	String	
评审的html地址。

pilot	Object	
评审所在的产品、项目或测试库。

description	String	
评审的描述。

participants	Object[]	
评审的关注人列表。

submitted_at	Number	
评审的提交时间。

submitted_by	Object	
评审的提交人。

completed_at	Number	
评审的完成时间。

completed_by	Object	
评审的完成人。

created_at	Number	
评审的创建时间。

created_by	Object	
评审的创建人。

updated_at	Number	
评审的更新时间。

updated_by	Object	
评审的更新人。

全量数据示例：
{
    "id": "6f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/6f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "completed",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
    "pilot": {
        "id": "63bb744314bd13c9def24cb4",
        "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "description": "这是一个评审的描述",
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
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
    "submitted_at": 1593290347,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "completed_at": 1593291347,
    "completed_by": {
        "id": "b2417f68e846aae315c85d24643678b0",
        "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
        "name": "mary",
        "display_name": "Mary",
        "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "updated_at": 1593291347,
    "updated_by": {
        "id": "b2417f68e846aae315c85d24643678b0",
        "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
        "name": "mary",
        "display_name": "Mary",
        "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
    }
}
引用数据示例：
{
    "id": "6f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/6f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "completed",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF"
}
创建一个评审
POST
https://rest_api_root/v1/reviews
权限: 企业令牌/用户令牌

请求参数
字段	类型	描述
title	String	
评审的标题。

pilot_id	String	
评审主体所在产品、项目或测试库的id。

principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

description可选	String	
评审的描述。

请求示例：
{
    "title": "这是一个评审",
    "pilot_id": "63bb744314bd13c9def24cb4",
    "principal_type": "idea",
    "description": "这是一个评审的描述"
}
响应示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/5f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "pending",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
    "pilot": {
        "id": "63bb744314bd13c9def24cb4",
        "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "description": "这是一个评审的描述",
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
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
    "submitted_at": null,
    "submitted_by": null,
    "completed_at": null,
    "completed_by": null,
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "updated_at": 1593291347,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    }
}
获取评审列表
GET
https://rest_api_root/v1/reviews?principal_type={principal_type}&pilot_id={pilot_id}
权限: 企业令牌/用户令牌

查询参数
字段	类型	描述
principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

pilot_id	String	
评审主体所在产品、项目或测试库的id。

status可选	String	
评审的状态。

允许值: pending, in_progress, completed, repealed

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5f168f764eba01a5278b87cd",
            "url": "https://rest_api_root/v1/reviews/5f168f764eba01a5278b87cd?principal_type=idea",
            "identifier": "SCR-R5",
            "title": "这是一个评审",
            "status": "completed",
            "principal_type": "idea",
            "short_id": "LsEy8mZF",
            "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
            "pilot": {
                "id": "63bb744314bd13c9def24cb4",
                "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
                "name": "示例产品",
                "identifier": "SLC",
                "is_archived": 0,
                "is_deleted": 0
            },
            "description": "这是一个评审的描述",
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
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
            "submitted_at": 1593290347,
            "submitted_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
            },
            "completed_at": 1593291347,
            "completed_by": {
                "id": "b2417f68e846aae315c85d24643678b0",
                "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
                "name": "mary",
                "display_name": "Mary",
                "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
            },
            "created_at": 1593290347,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
            },
            "updated_at": 1593291347,
            "updated_by": {
                "id": "b2417f68e846aae315c85d24643678b0",
                "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
                "name": "mary",
                "display_name": "Mary",
                "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
            }
        }
    ]
}
删除一个评审
DELETE
https://rest_api_root/v1/reviews/{review_id}?principal_type={principal_type}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
review_id	String	
评审的id。

查询参数
字段	类型	描述
principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

响应示例：
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/5f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "pending",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
    "pilot": {
        "id": "63bb744314bd13c9def24cb4",
        "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "description": "这是一个评审的描述",
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
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
    "submitted_at": null,
    "submitted_by": null,
    "completed_at": null,
    "completed_by": null,
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "updated_at": 1593291347,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    }
}
向评审中添加一个评审内容
POST
https://rest_api_root/v1/reviews/{review_id}/principals
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
review_id	String	
评审的id。

请求参数
字段	类型	描述
principal_id	String	
评审内容的id。

principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

请求示例：
{
    "principal_id": "63bb744514bd13c9def24ceb",
    "principal_type": "idea"
}
响应示例：
{
    "id": "63bb744514bd13c9def24ceb",
    "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
    "review": {
        "id": "68ccfe6b3eef8131da564e4a",
        "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
        "identifier": "SLC-R11",
        "title": "这是一个需求评审",
        "status": "pending",
        "principal_type": "idea",
        "short_id": "LK7Qy-NA",
        "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
    },
    "principal_type": "idea",
    "principal": {
        "id": "63bb744514bd13c9def24ceb",
        "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
        "identifier": "SLC-1",
        "title": "这是一个产品需求",
        "short_id": "Omi8PFL0",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
    }
}
获取评审内容列表
GET
https://rest_api_root/v1/reviews/{review_id}/principals?principal_type={principal_type}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
review_id	String	
评审的id。

查询参数
字段	类型	描述
principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744514bd13c9def24ceb",
            "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
            "review": {
                "id": "68ccfe6b3eef8131da564e4a",
                "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
                "identifier": "SLC-R11",
                "title": "这是一个需求评审",
                "status": "pending",
                "principal_type": "idea",
                "short_id": "LK7Qy-NA",
                "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
            },
            "principal_type": "idea",
            "principal": {
                "id": "63bb744514bd13c9def24ceb",
                "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
                "identifier": "SLC-1",
                "title": "这是一个产品需求",
                "short_id": "Omi8PFL0",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
            }
        }
    ]
}
在评审中移除一个评审内容
DELETE
https://rest_api_root/v1/reviews/{review_id}/principals/{principal_id}?principal_type={principal_type}
权限: 企业令牌/用户令牌

路径参数
字段	类型	描述
review_id	String	
评审的id。

principal_id	String	
评审内容的id。

查询参数
字段	类型	描述
principal_type	String	
评审主体的类型。

允许值: idea, work_item, test_case

响应示例：
{
    "id": "63bb744514bd13c9def24ceb",
    "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
    "review": {
        "id": "68ccfe6b3eef8131da564e4a",
        "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
        "identifier": "SLC-R11",
        "title": "这是一个需求评审",
        "status": "pending",
        "principal_type": "idea",
        "short_id": "LK7Qy-NA",
        "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
    },
    "principal_type": "idea",
    "principal": {
        "id": "63bb744514bd13c9def24ceb",
        "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
        "identifier": "SLC-1",
        "title": "这是一个产品需求",
        "short_id": "Omi8PFL0",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
    }
}