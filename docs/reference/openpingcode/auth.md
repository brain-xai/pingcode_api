
鉴权
客户端凭据
客户端凭据（OAuth2 Client Credentials）是最简单、最直接的授权方式，通过客户端凭据获取的访问令牌（access_token）不区分用户身份， 在PingCode中将这类访问令牌称为“企业令牌”，企业令牌具有系统管理员的权限，主要用于访问和操作全局的数据，请谨慎管理。 获取企业令牌时，需要提供client_id和client_secret， 您可以在PingCode企业后台的凭据管理中创建一个应用，配置数据范围，然后拿到client_id和client_secret。 使用企业令牌时，只需要在HTTP请求在headers中添加"authorization": "Bearer {access_token}"，即可访问受保护的数据。

获取企业令牌
access_token的有效期为30天，删除PingCode的应用或重置应用的Secret都会导致access_token失效。

GET
https://rest_api_root/v1/auth/token?grant_type=client_credentials
查询参数
字段	类型	描述
grant_type	String	
OAuth2的授予类型，这里固定为：client_credentials。

client_id	String	
PingCode应用的Client ID

client_secret	String	
PingCode应用的Secret

请求示例：
{
    "grant_type": "client_credentials",
    "client_id": "CVMTTGAwaYiz",
    "client_secret": "QjssecMkBwTzVDaNJsOIUvPp"
}
响应示例：
{
    "access_token": "e7321ca8-f724-4abd-9169-d76d095c6acf",
    "token_type": "Bearer",
    "expires_in": 1577808000
}
授权码
授权码（OAuth2 Authorization Code）是最常用的授权方式，主要用于企业员工管理自己的数据。 通过授权码获取的访问令牌（access_token）在PingCode中称为“用户令牌”，用户令牌属于某个用户所有，仅能访问这个用户有权限操作的数据。 第三方应用可以通过用户的手动授权获得该用户的用户令牌，然后通过用户令牌访问该用户有权限操作的数据。 在通过授权码的方式获取用户令牌时，需要提供client_id和code， 您可以在PingCode企业后台的凭据管理中创建一个应用，配置数据范围，然后拿到client_id和redirect_uri。 在您的第三方应用中引导用户访问请求授权页面，PingCode会询问该用户是否授权给您的应用。 得到用户许可后，浏览器会跳转redirect_uri页面，并在URL的参数中包含一个临时的code，然后您的应用可以根据client_id和code获取该用户的用户令牌。 使用用户令牌时，只需要在HTTP请求在headers中添加"authorization": "Bearer {access_token}"，即可访问受保护的数据。

请求授权
用户访问请求授权页面前，需要登录PingCode，否则页面自动跳转到PingCode的登录页面。需要注意，私有环境的授权页面地址为：https://xxxxxx/oauth2/authorize 。

GET
https://oauth2_root/authorize?response_type=code
查询参数
字段	类型	描述
response_type	String	
响应类型，这里固定为code类型。

client_id	String	
PingCode应用的Client ID

获取用户令牌
access_token的有效期为30天，删除PingCode的应用或重置应用的Secret都会导致access_token失效。

GET
https://rest_api_root/v1/auth/token?grant_type=authorization_code
查询参数
字段	类型	描述
grant_type	String	
OAuth2的授予类型，这里固定为：authorization_code。

client_id	String	
PingCode应用的Client ID

client_secret	String	
PingCode应用的Secret

code	String	
用户授权后，在回调地址中拿到的code值。

请求示例：
{
    "grant_type": "authorization_code",
    "client_id": "CVMTTGAwaYiz",
    "client_secret": "QjssecMkBwTzVDaNJsOIUvPp",
    "code": "9169-d76d095c6acf-f724-4abd-e7321ca8"
}
响应示例：
{
    "access_token": "e7321ca8-f724-4abd-9169-d76d095c6acf",
    "refresh_token": "f724-4abd-9169-e7321ca8-d76d095c6acf",
    "token_type": "Bearer",
    "expires_in": 1577808000
}
刷新用户令牌
通过refresh_token重新获取access_token，可以避免用户频繁重新授权。refresh_token的有效期为90天，删除PingCode的应用或重置应用的Secret都会导致refresh_token失效。

GET
https://rest_api_root/v1/auth/token?grant_type=refresh_token
查询参数
字段	类型	描述
grant_type	String	
OAuth2的授予类型，这里固定为：refresh_token。

refresh_token	String	
通过authorization_code获取access_token时，一起返回的refresh_token。

请求示例：
{
    "grant_type": "refresh_token",
    "refresh_token": "f724-4abd-9169-e7321ca8-d76d095c6acf"
}
响应示例：
{
    "access_token": "e7321ca8-f724-4abd-9169-d76d095c6acf",
    "token_type": "Bearer",
    "expires_in": 1577808000
}
