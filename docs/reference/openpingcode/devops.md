DevOps 数据集成
代码
托管平台
企业内实际的代码托管平台，例如Github或私有部署的Gitlab。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}

资源属性
字段	类型	描述
id	String	
托管平台的id。

url	String	
托管平台的资源地址。

name	String	
托管平台的名称。

type	String	
托管平台的类型。

允许值: github, gitlab, bitbucket, coding.net, gogs, git, svn, gerrit, other

description	String	
托管平台的描述。

全量数据示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
引用数据示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github"
}
创建一个托管平台
POST
https://rest_api_root/v1/scm/products
权限: 企业令牌

请求参数
字段	类型	描述
name	String	
代码托管平台的名称，在一个企业中这个名称是唯一的。

type	String	
代码托管平台的产品类型，主要用于显示图标。

允许值: github, gitlab, bitbucket, coding.net, gogs, git, svn, gerrit, other

description可选	String	
代码托管平台的简介

请求示例：
{
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
全量更新一个托管平台
PUT
https://rest_api_root/v1/scm/products/{product_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

请求参数
字段	类型	描述
name	String	
代码托管平台的名称，在一个企业中这个名称是唯一的。

type	String	
代码托管平台的产品类型，主要用于显示图标。

允许值: github, gitlab, bitbucket, coding.net, gogs, git, svn, gerrit, other

description可选	String	
代码托管平台简介。

请求示例：
{
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
部分更新一个托管平台
PATCH
https://rest_api_root/v1/scm/products/{product_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

请求参数
字段	类型	描述
name可选	String	
代码托管平台的名称，在一个企业中这个名称是唯一的。

type可选	String	
代码托管平台的产品类型，主要用于显示图标。

允许值: github, gitlab, bitbucket, coding.net, gogs, git, svn, gerrit, other

description可选	String	
代码托管平台简介。

请求示例：
{
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
获取托管平台列表
GET
https://rest_api_root/v1/scm/products
权限: 企业令牌

查询参数
字段	类型	描述
name可选	String	
代码托管平台的名称。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080765",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
            "name": "Github",
            "type": "github",
            "description": "Github公有云"
        }
    ]
}
托管平台用户
代码托管平台内实际的用户，用于在PingCode中显示该用户在代码托管平台上的名称、头像以及主页的信息。如果没有手动创建用户，在操作代码仓库、分支、拉取请求时，将自动创建仅包含该用户名的用户。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/users/{user_id}

资源属性
字段	类型	描述
id	String	
托管平台用户的id。

url	String	
托管平台用户的资源地址。

product	Object	
托管平台。

name	String	
托管平台用户的名称。

display_name	String	
托管平台用户的显示名。

html_url	String	
代码托管平台上的用户主页地址。

avatar_url	String	
代码托管平台上的用户头像地址。

全量数据示例：
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
引用数据示例：
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry"
}
创建一个托管平台用户
POST
https://rest_api_root/v1/scm/products/{product_id}/users
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

请求参数
字段	类型	描述
name	String	
代码托管平台上的用户名，同一代码托管平台下，该用户名是唯一的。

display_name可选	String	
用户显示名。

html_url可选	String	
代码托管平台上的用户主页地址。

avatar_url可选	String	
代码托管平台上的用户头像地址。

请求示例：
{
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
响应示例：
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
全量更新一个托管平台用户
PUT
https://rest_api_root/v1/scm/products/{product_id}/users/{user_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

user_id	String	
代码托管平台上的用户id。

请求参数
字段	类型	描述
name	String	
代码托管平台上的用户名，同一代码托管平台下，该用户名是唯一的。

display_name可选	String	
用户显示名。

html_url可选	String	
代码托管平台上的用户主页地址。

avatar_url可选	String	
代码托管平台上的用户头像地址。

请求示例：
{
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
响应示例：
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
部分更新一个托管平台用户
PATCH
https://rest_api_root/v1/scm/products/{product_id}/users/{user_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

user_id	String	
代码托管平台上的用户id。

请求参数
字段	类型	描述
name可选	String	
代码托管平台上的用户名，同一代码托管平台下，该用户名是唯一的。

display_name可选	String	
用户显示名。

html_url可选	String	
代码托管平台上的用户主页地址。

avatar_url可选	String	
代码托管平台上的用户头像地址。

请求示例：
{
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
响应示例：
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
获取托管平台用户列表
GET
https://rest_api_root/v1/scm/products/{product_id}/users
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

查询参数
字段	类型	描述
name可选	String	
代码托管平台上的用户名。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5666aea91f99e33cb7c44964",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "name": "terry",
            "display_name": "Terry",
            "html_url": "https://github.com/terrylee",
            "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
        }
    ]
}
代码仓库
代码托管平台内实际的代码仓库，用于在PingCode中显示代码仓库的详细信息。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}

资源属性
字段	类型	描述
id	String	
代码仓库的id。

url	String	
代码仓库的资源地址。

product	Object	
托管平台。

name	String	
代码仓库的名称。

full_name	String	
代码仓库的全称。

created_at	Number	
代码仓库的创建时间。

owner	Object	
代码仓库的拥有者。

is_fork	Boolean	
代码仓库是否是fork仓库。

is_private	Boolean	
代码仓库是否是私有仓库。

description	String	
代码仓库的描述。

html_url	String	
代码仓库的地址。

branches_url	String	
代码仓库的分支地址模板，链接后面括号里的值会被替换成真实地址。

commits_url	String	
代码仓库的提交地址模板，链接后面括号里的值会被替换成真实地址。

compare_url	String	
代码仓库的分支对比地址模板，链接后面括号里的值会被替换成真实地址。

pulls_url	String	
代码仓库的拉取请求地址模板，链接后面括号里的值会被替换成真实地址。

全量数据示例：
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
引用数据示例：
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919
}
创建一个代码仓库
POST
https://rest_api_root/v1/scm/products/{product_id}/repositories
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

请求参数
字段	类型	描述
name	String	
代码仓库的名称。

full_name	String	
代码仓库的全称。同一代码托管平台下，代码仓库的全称是唯一的。

description可选	String	
代码仓库的简介。

is_fork可选	Boolean	
是否是fork仓库。

is_private可选	Boolean	
是否是私有仓库。

owner_name可选	String	
代码仓库拥有者的用户名。

html_url可选	String	
代码仓库的地址。如果为空，在PingCode中不显示相关跳转链接。

branches_url可选	String	
代码仓库的分支地址，使用{branch}表示分支名。如果为空，在PingCode中不显示相关跳转链接。

commits_url可选	String	
代码仓库的提交地址，使用{sha}表示提交的SHA值。如果为空，在PingCode中不显示相关跳转链接。

compare_url可选	String	
代码仓库的分支对比地址，使用{base}和{head}表示基准分支名和需要进行对比的分支名。如果为空，在PingCode中不显示相关跳转链接。

pulls_url可选	String	
代码仓库的拉取请求地址，使用{number}表示拉取请求的编号。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "is_fork": false,
    "is_private": false,
    "owner_name": "terry",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
响应示例：
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
全量更新一个代码仓库
PUT
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

请求参数
字段	类型	描述
name	String	
代码仓库的名称。

full_name	String	
代码仓库的全称。同一代码托管平台下，代码仓库的全称是唯一的。

description可选	String	
代码仓库的简介。

is_fork可选	Boolean	
是否是fork仓库。

is_private可选	Boolean	
是否是私有仓库。

owner_name可选	String	
代码仓库拥有者的用户名。

html_url可选	String	
代码仓库在代码托管平台上的地址。如果为空，在PingCode中不显示相关跳转链接。

branches_url可选	String	
代码仓库的分支地址，使用{branch}表示分支名。如果为空，在PingCode中不显示相关跳转链接。

commits_url可选	String	
代码仓库的提交地址，使用{sha}表示提交的SHA值。如果为空，在PingCode中不显示相关跳转链接。

compare_url可选	String	
代码仓库的分支对比地址，使用{base}和{head}表示基准分支名和需要进行对比的分支名。如果为空，在PingCode中不显示相关跳转链接。

pulls_url可选	String	
代码仓库的拉取请求地址，使用{number}表示拉取请求的编号。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "is_fork": false,
    "is_private": false,
    "owner_name": "terry",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
响应示例：
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
部分更新一个代码仓库
PATCH
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

请求参数
字段	类型	描述
name可选	String	
代码仓库的名称。

full_name可选	String	
代码仓库的全称。同一代码托管平台下，代码仓库的全称是唯一的。

description可选	String	
代码仓库的简介。

is_fork可选	Boolean	
是否是fork仓库。

is_private可选	Boolean	
是否是私有仓库。

owner_name可选	String	
代码仓库拥有者的用户名。

html_url可选	String	
代码仓库在代码托管平台上的地址。如果为空，在PingCode中不显示相关跳转链接。

branches_url可选	String	
代码仓库的分支地址，使用{branch}表示分支名。如果为空，在PingCode中不显示相关跳转链接。

commits_url可选	String	
代码仓库的提交地址，使用{sha}表示提交的SHA值。如果为空，在PingCode中不显示相关跳转链接。

compare_url可选	String	
代码仓库的分支对比地址，使用{base}和{head}表示基准分支名和需要进行对比的分支名。如果为空，在PingCode中不显示相关跳转链接。

pulls_url可选	String	
代码仓库的拉取请求地址，使用{number}表示拉取请求的编号。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "is_fork": false,
    "is_private": false,
    "owner_name": "terry",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
响应示例：
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
获取代码仓库列表
GET
https://rest_api_root/v1/scm/products/{product_id}/repositories
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

查询参数
字段	类型	描述
full_name可选	String	
代码仓库的全称。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080766",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "name": "ngx-planet",
            "full_name": "worktile/ngx-planet",
            "created_at": 1403018919,
            "owner": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
            },
            "is_fork": false,
            "is_private": false,
            "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
            "html_url": "https://github.com/worktile/ngx-planet",
            "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
            "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
            "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
            "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
        }
    ]
}
代码分支
代码仓库内实际的分支，用于在PingCode中显示分支的详细信息。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/branches/{branch_id}

资源属性
字段	类型	描述
id	String	
代码分支的id。

url	String	
代码分支的资源地址。

product	Object	
代码分支的托管平台。

repository	Object	
代码分支的代码仓库。

name	String	
代码分支的名称。

created_at	Number	
代码分支的创建时间。

sender	Object	
代码分支的创建者。

is_default	Boolean	
代码分支是否为默认分支。

work_items	Object[]	
代码分支关联的工作项列表。

全量数据示例：
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": false,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
引用数据示例：
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "name": "terry/#PLM-001",
    "created_at": 1403018919
}
创建一个代码分支
POST
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/branches
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

请求参数
字段	类型	描述
name	String	
分支的名称。同一代码仓库下，分支的名称是唯一的。

sender_name	String	
分支创建者的用户名。

is_default可选	Boolean	
是否设置为默认分支。当创建分支时，如果当前仓库的分支数为0，则新创建的分支会自动设置为该仓库的默认分支。如果创建分支时设置了该分支为默认分支，并且仓库已经有默认分支存在，那么其他分支将被取消默认属性，而该分支将被设置为新的默认分支。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将分支与一个或多个工作项关联，分支和工作项关联后，分支下所有的提交都会和该工作项关联。

请求示例：
{
    "name": "terry/#PLM-001",
    "sender_name": "terry",
    "is_default": true,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": true,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
部分更新一个代码分支
PATCH
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/branches/{branch_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

branch_id	String	
分支的id。

请求参数
字段	类型	描述
is_default可选	Boolean	
是否设置为默认分支。该值只能是true，设置该分支为默认分支后将取消其他分支的默认属性。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将分支与一个或多个工作项关联，分支和工作项关联后，分支下所有的提交都会和该工作项关联。

请求示例：
{
    "is_default": true,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": true,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
获取代码分支列表
GET
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/branches
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

查询参数
字段	类型	描述
name可选	String	
分支的名称。

work_item_id可选	String	
工作项的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080767",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
            },
            "name": "terry/#PLM-001",
            "created_at": 1403018919,
            "sender": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
            },
            "is_default": false,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
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
            ]
        }
    ]
}
删除一个代码分支
删除分支后，不会移除该分支和工作项的关联历史，在关联历史中，依然可以查询到删除的分支。默认分支不能被删除。

DELETE
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/branches/{branch_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

branch_id	String	
分支的id。

响应示例：
{
    "id": "564587fe700d43b81b080768",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080768",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": false,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
提交
实际的代码提交记录，用于在PingCode中显示提交的详细信息。提交并不会自动和代码仓库关联，需要通过提交引用与之关联。
资源地址： GET https://rest_api_root/v1/scm/commits/{commit_id_or_sha}

资源属性
字段	类型	描述
id	String	
提交的id。

url	String	
提交的资源地址。

sha	String	
提交的SHA值。

message	String	
提交的描述信息。

committer_name	String	
提交者的用户名。

committed_at	Number	
提交的时间。

tree_id	String	
提交树的SHA值。

files_added	String[]	
提交新增文件的文件名列表。

files_removed	String[]	
提交移除文件的文件名列表。

files_modified	String[]	
提交更新文件的文件名列表。

file_changed_count	Number	
提交更新文件数量。

work_items	Object[]	
提交关联的PingCode的工作项列表。

全量数据示例：
{
    "id": "5e3bb2128cfda459bbafa3fb",
    "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919,
    "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
    "files_added": [
        "index.ts"
    ],
    "files_removed": [
        "utilities.ts"
    ],
    "files_modified": [
        "README.md"
    ],
    "file_changed_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
引用数据示例：
{
    "id": "5e3bb2128cfda459bbafa3fb",
    "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919
}
创建一个提交
POST
https://rest_api_root/v1/scm/commits
权限: 企业令牌

请求参数
字段	类型	描述
sha	String	
提交的SHA值。

message	String	
提交的描述信息。

committer_name	String	
提交者的用户名。

committed_at	Number	
提交的时间。

tree_id可选	String	
提交树的SHA值。

files_added	String[]	
新增文件的文件名列表。

files_removed	String[]	
移除文件的文件名列表。

files_modified	String[]	
更新文件的文件名列表。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将提交与一个或多个工作项关联。

请求示例：
{
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919,
    "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
    "files_added": [
        "index.ts"
    ],
    "files_removed": [
        "utilities.ts"
    ],
    "files_modified": [
        "README.md"
    ],
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "5e3bb2128cfda459bbafa3fb",
    "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919,
    "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
    "files_added": [
        "index.ts"
    ],
    "files_removed": [
        "utilities.ts"
    ],
    "files_modified": [
        "README.md"
    ],
    "file_changed_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
获取提交列表
GET
https://rest_api_root/v1/scm/commits
权限: 企业令牌

查询参数
字段	类型	描述
sha可选	String	
提交的SHA值。

work_item_id可选	String	
工作项的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e3bb2128cfda459bbafa3fb",
            "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
            "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
            "message": "feat(scope): #PLM-001 initialization code structure",
            "committer_name": "terry",
            "committed_at": 1403018919,
            "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
            "files_added": [
                "index.ts"
            ],
            "files_removed": [
                "utilities.ts"
            ],
            "files_modified": [
                "README.md"
            ],
            "file_changed_count": 3,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
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
            ]
        }
    ]
}
提交引用
提交引用是提交与分支的一种关联关系，一个提交可以与多个分支关联，一个分支也可以与多个提交关联。 当提交与分支关联后，提交会自动与由此分支发起的拉取请求关联，当拉取请求合并后，这些关联的提交将自动被标记为“已合并”状态。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/refs/{ref_id}

资源属性
字段	类型	描述
id	String	
提交引用的id。

url	String	
提交引用的资源地址。

product	Object	
提交引用的托管平台。

repository	Object	
提交引用的代码仓库。

commit	Object	
提交引用的代码提交。

meta	Object	
提交引用的实体，如分支。

全量数据示例：
{
    "id": "5e451b7dd704c212f7de8b4f",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "commit": {
        "id": "5e3bb2128cfda459bbafa3fb",
        "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
        "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
        "message": "feat(scope): #PLM-001 initialization code structure",
        "committer_name": "terry",
        "committed_at": 1403018919
    },
    "meta": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "type": "branch"
    }
}
引用数据示例：
{
    "id": "5e451b7dd704c212f7de8b4f",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f"
}
创建一个提交引用
POST
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/refs
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

请求参数
字段	类型	描述
sha	String	
提交的SHA值。

meta_type	String	
引用实体的类型。

允许值: branch

meta_id	String	
引用实体的id，例如：branch_id。

请求示例：
{
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "meta_type": "branch",
    "meta_id": "564587fe700d43b81b080767"
}
响应示例：
{
    "id": "5e451b7dd704c212f7de8b4f",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "commit": {
        "id": "5e3bb2128cfda459bbafa3fb",
        "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
        "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
        "message": "feat(scope): #PLM-001 initialization code structure",
        "committer_name": "terry",
        "committed_at": 1403018919
    },
    "meta": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "type": "branch"
    }
}
获取提交引用列表
GET
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/refs
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

查询参数
字段	类型	描述
meta_type	String	
引用实体的类型。

允许值: branch

meta_id	String	
引用实体的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e451b7dd704c212f7de8b4f",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
            },
            "commit": {
                "id": "5e3bb2128cfda459bbafa3fb",
                "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
                "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
                "message": "feat(scope): #PLM-001 initialization code structure",
                "committer_name": "terry",
                "committed_at": 1403018919
            },
            "meta": {
                "id": "564587fe700d43b81b080767",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
                "name": "terry/#PLM-001",
                "type": "branch"
            }
        }
    ]
}
拉取请求
代码仓库内实际的拉取请求，用于在PingCode中显示拉取请求的详细信息。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}

资源属性
字段	类型	描述
id	String	
拉取请求的id。

url	String	
拉取请求的资源地址。

product	Object	
拉取请求的托管平台。

repository	Object	
拉取请求的代码仓库。

title	String	
拉取请求的标题。

number	Number	
拉取请求的编号。

status	String	
拉取请求的状态。

允许值: open, closed, merged, abandoned

description	String	
拉取请求的描述。

author	Object	
拉取请求的创建者。

source_branch	Object	
拉取请求的源分支。

target_branch	Object	
拉取请求的目标分支。

created_at	Number	
拉取请求的创建时间。

merged_at	Number	
拉取请求的合并的时间。

merged_commit_sha	String	
拉取请求的源分支最后一次提交的SHA值。

merged_by	Object	
拉取请求的合并者。

comments_count	Number	
拉取请求的评论数量。

review_comments_count	Number	
拉取请求的代码评审评论数量。

commits_count	Number	
拉取请求的提交数量。

additions_count	Number	
拉取请求的新增文件数量。

deletions_count	Number	
拉取请求的删除文件数量。

changed_files_count	Number	
拉取请求的更改文件数量。

work_items	Object[]	
拉取请求关联的PingCode的工作项列表。

全量数据示例：
{
    "id": "594587fe700d43b81b080789",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "title": "fix(doc): #PLM-001 fix document title",
    "number": 1,
    "status": "merged",
    "description": "Please give some great suggestions",
    "author": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "source_branch": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "create_at": 1403018919
    },
    "target_branch": {
        "id": "564587fe700d43b81b080776",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
        "name": "develop",
        "create_at": 1402018919
    },
    "created_at": 1403014000,
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
引用数据示例：
{
    "id": "594587fe700d43b81b080789",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
    "number": 1
}
创建一个拉取请求
POST
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

请求参数
字段	类型	描述
title	String	
拉取请求的标题。

number	Number	
拉取请求的编号。同一代码仓库下，该值是唯一的。

creator_name	String	
拉取请求创建者的用户名。

source_branch_id可选	String	
源分支的id。

target_branch_id	String	
目标分支的id。

status	String	
拉取请求的状态。

允许值: open, closed, merged, abandoned

description可选	String	
拉取请求的描述。

merged_at可选	Number	
拉取请求合并的时间。当拉取请求状态为merged时，该值为必填。

merged_commit_sha可选	String	
源分支最后一次提交的SHA值。当拉取请求状态为merged时，该值为必填。

merged_by_name可选	String	
拉取请求合并者的用户名。当拉取请求状态为merged时，该值为必填。

comments_count可选	Number	
拉取请求的评论数量。

review_comments_count可选	Number	
代码评审的评论数量。

commits_count可选	Number	
代码的提交数量。

additions_count可选	Number	
新增文件的数量。

deletions_count可选	Number	
删除文件的数量。

changed_files_count可选	Number	
更改文件的数量。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将拉取请求与一个或多个工作项关联。

请求示例：
{
    "title": "fix(doc): #PLM-001 fix document title",
    "number": 1,
    "creator_name": "terry",
    "description": "Please give some great suggestions",
    "source_branch_id": "564587fe700d43b81b080767",
    "target_branch_id": "564587fe700d43b81b080776",
    "status": "merged",
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by_name": "terry",
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "594587fe700d43b81b080789",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "title": "fix(doc): #PLM-001 fix document title",
    "number": 1,
    "status": "merged",
    "description": "Please give some great suggestions",
    "author": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "source_branch": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "create_at": 1403018919
    },
    "target_branch": {
        "id": "564587fe700d43b81b080776",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
        "name": "develop",
        "create_at": 1402018919
    },
    "created_at": 1463014000,
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
全量更新一个拉取请求
PUT
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

pull_request_id	String	
拉取请求的id。

请求参数
字段	类型	描述
title	String	
拉取请求的标题。

creator_name	String	
拉取请求创建者的用户名。

source_branch_id	String	
源分支的id。

target_branch_id	String	
目标分支的id。

status	String	
拉取请求的状态。

允许值: open, closed, merged, abandoned

description可选	String	
拉取请求的描述。

merged_at可选	Number	
拉取请求合并的时间。当拉取请求状态为merged时，该值为必填。

merged_commit_sha可选	String	
源分支最后一次提交的SHA值。当拉取请求状态为merged时，该值为必填。

merged_by_name可选	String	
拉取请求合并者的用户名。当拉取请求状态为merged时，该值为必填。

comments_count可选	Number	
拉取请求的评论数量。

review_comments_count可选	Number	
代码评审的评论数量。

commits_count可选	Number	
代码的提交数量。

additions_count可选	Number	
新增文件的数量。

deletions_count可选	Number	
删除文件的数量。

changed_files_count可选	Number	
更改文件的数量。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将拉取请求与一个或多个工作项关联。

请求示例：
{
    "title": "fix(doc): #PLM-001 fix document title",
    "number": 1,
    "creator_name": "terry",
    "description": "Please give some great suggestions",
    "source_branch_id": "564587fe700d43b81b080767",
    "target_branch_id": "564587fe700d43b81b080776",
    "status": "merged",
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by_name": "terry",
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "594587fe700d43b81b080789",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "title": "fix(doc): #PLM-001 fix document title",
    "number": 1,
    "status": "merged",
    "description": "Please give some great suggestions",
    "author": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "source_branch": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "create_at": 1403018919
    },
    "target_branch": {
        "id": "564587fe700d43b81b080776",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
        "name": "develop",
        "create_at": 1402018919
    },
    "created_at": 1403014000,
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by": {
        "id": "5666aea91f99e33cb7c44968",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
部分更新一个拉取请求
PATCH
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

pull_request_id	String	
拉取请求的id。

请求参数
字段	类型	描述
status	String	
拉取请求的状态。

允许值: open, closed, merged, abandoned

title可选	String	
拉取请求的标题。

creator_name可选	String	
拉取请求创建者的用户名。

description可选	String	
拉取请求的描述。

target_branch_id可选	String	
目标分支的id。

source_branch_id可选	String	
源分支的id。

merged_at可选	Number	
拉取请求合并的时间。当拉取请求状态为merged时，该值为必填。

merged_commit_sha可选	String	
源分支最后一次提交的SHA值。当拉取请求状态为merged时，该值为必填。

merged_by_name可选	String	
拉取请求合并者的用户名。当拉取请求状态为merged时，该值为必填。

comments_count可选	Number	
拉取请求的评论数量。

review_comments_count可选	Number	
代码评审的评论数量。

commits_count可选	Number	
代码的提交数量。

additions_count可选	Number	
新增文件的数量。

deletions_count可选	Number	
删除文件的数量。

changed_files_count可选	Number	
更改文件的数量。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将拉取请求与一个或多个工作项关联。

请求示例：
{
    "title": "fix(doc): #PLM-001 fix document title",
    "description": "Please give some great suggestions",
    "status": "merged",
    "target_branch_id": "564587fe700d43b81b080776",
    "source_branch_id": "564587fe700d43b81b080767",
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by_name": "terry",
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "594587fe700d43b81b080789",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "title": "fix(doc): #PLM-001 fix document title",
    "number": 1,
    "status": "merged",
    "description": "Please give some great suggestions",
    "author": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "source_branch": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "create_at": 1403018919
    },
    "target_branch": {
        "id": "564587fe700d43b81b080776",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
        "name": "develop",
        "create_at": 1402018919
    },
    "created_at": 1403014000,
    "merged_at": 1473018919,
    "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "merged_by": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "comments_count": 2,
    "review_comments_count": 2,
    "commits_count": 2,
    "additions_count": 0,
    "deletions_count": 0,
    "changed_files_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
获取拉取请求列表
GET
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

查询参数
字段	类型	描述
number可选	number	
拉取请求的编号。

work_item_id可选	String	
工作项的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "594587fe700d43b81b080789",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
            },
            "title": "fix(doc): #PLM-001 fix document title",
            "number": 1,
            "status": "merged",
            "description": "Please give some great suggestions",
            "author": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
            },
            "source_branch": {
                "id": "564587fe700d43b81b080767",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
                "name": "terry/#PLM-001",
                "create_at": 1403018919
            },
            "target_branch": {
                "id": "564587fe700d43b81b080776",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
                "name": "develop",
                "create_at": 1402018919
            },
            "created_at": 1403014000,
            "merged_at": 1473018919,
            "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
            "merged_by": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
            },
            "comments_count": 2,
            "review_comments_count": 2,
            "commits_count": 2,
            "additions_count": 0,
            "deletions_count": 0,
            "changed_files_count": 3,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
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
            ]
        }
    ]
}
代码评审
拉取请求实际的代码评审记录，用于在PingCode中显示代码评审的详细信息。
资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews/{review_id}

资源属性
字段	类型	描述
id	String	
代码评审的id。

url	String	
代码评审的资源地址。

product	Object	
代码评审的托管平台。

repository	Object	
代码评审的代码仓库。

pull_request	Object	
代码评审的拉取请求。

reviewer	Object	
代码评审的评审人。

status	String	
代码评审的状态。

允许值: comment, approved, request_changes

description	String	
代码评审的描述。

submitted_at	Number	
代码评审的提交时间。

html_url	String	
代码评审的地址。

全量数据示例：
{
    "id": "524587fe700d43b81b080988",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "pull_request": {
        "id": "594587fe700d43b81b080789",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
        "number": 1
    },
    "reviewer": {
        "id": "5999aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
        "name": "anytao"
    },
    "status": "approved",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
引用数据示例：
{
  "id": "524587fe700d43b81b080988",
  "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
}
创建一个代码评审
POST
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

pull_request_id	String	
拉取请求的id。

请求参数
字段	类型	描述
status	String	
代码评审的状态。

允许值: comment, approved, request_changes

reviewer_name	String	
评审人的用户名。

description可选	String	
代码评审的描述。

submitted_at	Number	
提交的时间。

html_url可选	String	
代码评审的地址。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "status": "approved",
    "reviewer_name": "anytao",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
响应示例：
{
    "id": "524587fe700d43b81b080988",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "pull_request": {
        "id": "594587fe700d43b81b080789",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
        "number": 1
    },
    "reviewer": {
        "id": "5999aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
        "name": "anytao"
    },
    "status": "approved",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
全量更新一个代码评审
PUT
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews/{review_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

pull_request_id	String	
拉取请求的id。

review_id	String	
代码评审的id。

请求参数
字段	类型	描述
reviewer_name	String	
评审人的用户名。

status	String	
代码评审的状态。

允许值: comment, approved, request_changes

submitted_at	Number	
提交的时间。

description可选	String	
代码评审的描述。

html_url可选	String	
代码评审的地址。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "reviewer_name": "anytao",
    "status": "approved",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
响应示例：
{
    "id": "524587fe700d43b81b080988",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "pull_request": {
        "id": "594587fe700d43b81b080789",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
        "number": 1
    },
    "reviewer": {
        "id": "5999aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
        "name": "anytao"
    },
    "status": "approved",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
部分更新一个代码评审
PATCH
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews/{review_id}
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

pull_request_id	String	
拉取请求的id。

review_id	String	
代码评审的id。

请求参数
字段	类型	描述
reviewer_name可选	String	
评审人的用户名。

status可选	String	
代码评审的状态。

允许值: comment, approved, request_changes

description可选	String	
代码评审的描述。

submitted_at可选	Number	
提交的时间。

html_url可选	String	
代码评审的地址。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "reviewer_name": "anytao",
    "status": "approved",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
响应示例：
{
    "id": "524587fe700d43b81b080988",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "pull_request": {
        "id": "594587fe700d43b81b080789",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
        "number": 1
    },
    "reviewer": {
        "id": "5999aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
        "name": "anytao"
    },
    "status": "approved",
    "description": "Review has approved",
    "submitted_at": 1403014111,
    "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
获取代码评审列表
GET
https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews
权限: 企业令牌

路径参数
字段	类型	描述
product_id	String	
代码托管平台的id。

repository_id	String	
代码仓库的id。

pull_request_id	String	
拉取请求的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "524587fe700d43b81b080988",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
            },
            "pull_request": {
                "id": "594587fe700d43b81b080789",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
                "number": 1
            },
            "reviewer": {
                "id": "5999aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
                "name": "anytao"
            },
            "status": "approved",
            "description": "Review has approved",
            "submitted_at": 1403014111,
            "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
        }
    ]
}
构建
构建记录
企业内实际的构建记录，用于在PingCode中显示构建的详细信息。
资源地址： GET https://rest_api_root/v1/build/builds/{build_id}

资源属性
字段	类型	描述
id	String	
构建记录的id。

url	String	
构建记录的资源地址。

identifier	String	
构建记录的编号。

name	String	
构建记录的名称。

job_url	String	
构建任务的地址。

provider	String	
构建工具的名称。

允许值: bamboo, bitbucket, jenkins, other

result_overview	String	
构建结果的概述。

result_url	String	
构建结果的地址。

start_at	Number	
构建开始的时间。

end_at	Number	
构建结束的时间。

status	String	
构建的状态。

允许值: success, failure, unknown

duration	Number	
构建持续的时间，单位为秒。

work_items	Object[]	
构建关联的PingCode的工作项列表。

全量数据示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
    "identifier": "131",
    "name": "unit-test",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "status": "success",
    "end_at": 1583290347,
    "duration": 38,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
引用数据示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
    "identifier": "131",
    "name": "unit-test",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "status": "success",
    "end_at": 1583290347,
    "duration": 38
}
创建一条构建记录
POST
https://rest_api_root/v1/build/builds
权限: 企业令牌

请求参数
字段	类型	描述
name	String	
构建的名称。

identifier	String	
构建的编号。

job_url可选	String	
构建任务的地址。如果为空，在PingCode中不显示相关跳转链接。

provider	String	
构建工具的名称。

允许值: bamboo, bitbucket, jenkins, other

result_overview可选	String	
构建结果的概述。

result_url可选	String	
构建结果的地址。如果为空，在PingCode中不显示相关的跳转链接。

start_at	Number	
构建开始的时间。

end_at	Number	
构建结束的时间。

duration	Number	
构建持续的时间。单位为秒。

status	String	
构建的状态。

允许值: success, failure

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将构建与一个或多个工作项关联。

请求示例：
{
    "name": "unit-test",
    "identifier": "131",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "duration": 38,
    "status": "success",
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
    "name": "unit-test",
    "identifier": "131",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "status": "success",
    "end_at": 1583290347,
    "duration": 38,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
全量更新一条构建记录
PUT
https://rest_api_root/v1/build/builds/{build_id}
权限: 企业令牌

路径参数
字段	类型	描述
build_id	String	
构建的id。

请求参数
字段	类型	描述
name	String	
构建的名称。

identifier	String	
构建的编号。

job_url可选	String	
构建任务的地址。如果为空，在PingCode中不显示相关跳转链接。

provider	String	
构建工具的名称。

允许值: bamboo, bitbucket, jenkins, other

result_overview可选	String	
构建结果的概述。

result_url可选	String	
构建结果的地址。如果为空，在PingCode中不显示相关的跳转链接。

start_at	Number	
构建开始的时间。

end_at	Number	
构建结束的时间。

duration	Number	
构建持续的时间。单位为秒。

status	String	
构建的状态。

允许值: success, failure

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将构建与一个或多个工作项关联。

请求示例：
{
    "name": "unit-test",
    "identifier": "131",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "duration": 38,
    "status": "success",
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
    "identifier": "131",
    "name": "unit-test",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "status": "success",
    "end_at": 1583290347,
    "duration": 38,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
部分更新一条构建记录
PATCH
https://rest_api_root/v1/build/builds/{build_id}
权限: 企业令牌

路径参数
字段	类型	描述
build_id	String	
构建的id。

请求参数
字段	类型	描述
name可选	String	
构建的名称。

identifier可选	String	
构建的编号。

job_url可选	String	
构建任务的地址。如果为空，在PingCode中不显示相关跳转链接。

provider可选	String	
构建工具的名称。

允许值: bamboo, bitbucket, jenkins, other

result_overview可选	String	
构建结果的概述。

result_url可选	String	
构建结果的地址。如果为空，在PingCode中不显示相关的跳转链接。

start_at可选	Number	
构建开始的时间。

end_at可选	Number	
构建结束的时间。

status可选	String	
构建的状态。

允许值: success, failure

duration可选	Number	
构建持续的时间。单位为秒。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将构建与一个或多个工作项关联。

请求示例：
{
    "name": "unit-test",
    "identifier": "131",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "status": "success",
    "duration": 38,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
    "identifier": "131",
    "name": "unit-test",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "status": "failure",
    "end_at": 1583290347,
    "duration": 38,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
获取构建记录列表
GET
https://rest_api_root/v1/build/builds
权限: 企业令牌

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080765",
            "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
            "identifier": "131",
            "name": "unit-test",
            "job_url": "https://your-job-url",
            "provider": "jenkins",
            "result_overview": "1000 test cases pass",
            "result_url": "https://your-result-url",
            "start_at": 1583290309,
            "status": "success",
            "end_at": 1583290347,
            "duration": 38,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
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
            ]
        }
    ]
}
删除一条构建记录
DEL
https://rest_api_root/v1/build/builds/{build_id}
权限: 企业令牌

路径参数
字段	类型	描述
build_id	String	
构建的id。

响应示例：
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
    "identifier": "131",
    "name": "unit-test",
    "job_url": "https://your-job-url",
    "provider": "jenkins",
    "result_overview": "1000 test cases pass",
    "result_url": "https://your-result-url",
    "start_at": 1583290309,
    "status": "success",
    "end_at": 1583290347,
    "duration": 38,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
交付
环境
企业内实际的部署环境，用于在PingCode中显示各个环境的部署信息。
资源地址： GET https://rest_api_root/v1/release/environments/{env_id}

资源属性
字段	类型	描述
id	String	
环境的id。

url	String	
环境的资源地址。

name	String	
环境的名称。

html_url	String	
环境的真实地址。

全量数据示例：
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
引用数据示例：
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production"
}
创建一个环境
POST
https://rest_api_root/v1/release/environments
权限: 企业令牌

请求参数
字段	类型	描述
name	String	
环境的名称。同一团队内，环境的名称是唯一的。

html_url可选	String	
环境的地址。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "name": "Production",
    "html_url": "https://your-environment-url"
}
响应示例：
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
全量更新一个环境
PUT
https://rest_api_root/v1/release/environments/{env_id}
权限: 企业令牌

路径参数
字段	类型	描述
env_id	String	
环境的id。

请求参数
字段	类型	描述
name	String	
环境的名称。同一团队内，环境的名称是唯一的。

html_url可选	String	
环境的地址。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "name": "Production",
    "html_url": "https://your-environment-url"
}
响应示例：
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
部分更新一个环境
PATCH
https://rest_api_root/v1/release/environments/{env_id}
权限: 企业令牌

路径参数
字段	类型	描述
env_id	String	
环境的id。

请求参数
字段	类型	描述
name可选	String	
环境的名称。同一团队内，环境的名称是唯一的。

html_url可选	String	
环境的地址。如果为空，在PingCode中不显示相关跳转链接。

请求示例：
{
    "name": "Production",
    "html_url": "https://your-environment-url"
}
响应示例：
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
获取环境列表
GET
https://rest_api_root/v1/release/environments
权限: 企业令牌

查询参数
字段	类型	描述
name	String	
环境的名称。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080123",
            "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
            "name": "Production",
            "html_url": "https://your-environment-url"
        }
    ]
}
删除一个环境
删除环境之前，需要先删除与该环境关联的部署。

DELETE
https://rest_api_root/v1/release/environments/{env_id}
权限: 企业令牌

路径参数
字段	类型	描述
env_id	String	
环境的id。

响应示例：
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
部署
企业内实际的部署记录，用于在PingCode中显示部署的详细信息。
资源地址： GET https://rest_api_root/v1/release/deploys/{deploy_id}

资源属性
字段	类型	描述
id	String	
部署的id。

url	String	
部署的资源地址。

status	String	
部署的状态。

允许值: not_deployed, deployed

release_name	String	
发布的名称。

environment	Object	
发布的环境。

release_url	String	
版本发布的地址。

start_at	Number	
部署开始的时间。

end_at	Number	
部署结束的时间。

duration	Number	
部署持续的时间。单位为秒。

work_items	Object[]	
部署关联的PingCode的工作项列表。

全量数据示例：
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
引用数据示例：
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0"
}
创建一个部署
POST
https://rest_api_root/v1/release/deploys
权限: 企业令牌

请求参数
字段	类型	描述
status	String	
部署的状态。

允许值: not_deployed, deployed

env_id	String	
环境的id。

release_name	String	
发布的名称。

release_url可选	String	
版本发布的地址。如果为空，在PingCode中不显示相关跳转链接。

start_at	Number	
部署开始的时间。

end_at	Number	
部署结束的时间。

duration	Number	
部署持续的时间。单位为秒。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将部署与一个或多个工作项关联。

请求示例：
{
    "status": "deployed",
    "env_id": "564587fe700d43b81b080123",
    "release_name": "1.1.0",
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
全量更新一个部署
PUT
https://rest_api_root/v1/release/deploys/{deploy_id}
权限: 企业令牌

路径参数
字段	类型	描述
deploy_id	String	
部署的id。

请求参数
字段	类型	描述
status	String	
部署的状态。

允许值: not_deployed, deployed

env_id	String	
环境的id。

release_name	String	
版本发布的名称。

release_url可选	String	
版本发布的地址。如果为空，在PingCode中不显示相关跳转链接。

start_at	Number	
部署开始的时间。

end_at	Number	
部署结束的时间。

duration	Number	
部署持续的时间。单位为秒。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将部署与一个或多个工作项关联。

请求示例：
{
    "status": "deployed",
    "env_id": "564587fe700d43b81b080123",
    "release_name": "1.1.0",
    "release_url": "https://your-release-host/release",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://your-release-host/release",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
部分更新一个部署
PATCH
https://rest_api_root/v1/release/deploys/{deploy_id}
权限: 企业令牌

路径参数
字段	类型	描述
deploy_id	String	
部署的id。

请求参数
字段	类型	描述
status可选	String	
部署的状态。

允许值: not_deployed, deployed

env_id可选	String	
环境的id。

release_name可选	String	
版本发布的名称。

release_url可选	String	
版本发布的地址。如果为空，在PingCode中不显示相关跳转链接。

start_at可选	Number	
部署开始的时间。

end_at可选	Number	
部署结束的时间。

duration可选	Number	
部署持续的时间。单位为秒。

work_item_identifiers可选	String[]	
PingCode工作项编号的列表。通过该参数将部署与一个或多个工作项关联。

请求示例：
{
    "status": "deployed",
    "env_id": "564587fe700d43b81b080123",
    "release_name": "1.1.0",
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_item_identifiers": [
        "PLM-001"
    ]
}
响应示例：
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}
获取部署列表
GET
https://rest_api_root/v1/release/deploys
权限: 企业令牌

查询参数
字段	类型	描述
env_id可选	String	
环境的id。

响应示例：
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080339",
            "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
            "status": "deployed",
            "release_name": "1.1.0",
            "environment": {
                "id": "564587fe700d43b81b080123",
                "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
                "name": "Production"
            },
            "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
            "start_at": 1583143467,
            "end_at": 1583143667,
            "duration": 200,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
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
            ]
        }
    ]
}
删除一个部署
DELETE
https://rest_api_root/v1/release/deploys/{deploy_id}
权限: 企业令牌

路径参数
字段	类型	描述
deploy_id	String	
部署的id。

响应示例：
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
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
    ]
}