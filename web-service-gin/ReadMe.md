REST（英文：Representational State Transfer，简称REST）

描述了一个架构样式的网络系统，比如 web 应用程序。它首次出现在 2000 年 Roy Fielding 的博士论文中，Roy Fielding是 HTTP 规范的主要编写者之一。在目前主流的三种Web服务交互方案中，REST相比于SOAP（Simple Object Access protocol，简单对象访问协议）以及XML-RPC更加简单明了，无论是对URL的处理还是对Payload的编码，REST都倾向于用更加简单轻量的方法设计和实现。值得注意的是REST并没有一个明确的标准，而更像是一种设计的风格。

RESTFUL特点包括：

1、每一个URI代表1种资源；
2、客户端使用GET、POST、PUT、DELETE4个表示操作方式的动词对服务端资源进行操作：GET用来获取资源，POST用来新建资源（也可以用于更新资源），PUT用来更新资源，DELETE用来删除资源；
3、通过操作资源的表现形式来操作资源；
4、资源的表现形式是XML或者HTML；
5、客户端与服务端之间的交互在请求之间是无状态的，从客户端到服务端的每个请求都必须包含理解请求所必需的信息。


RESTful与 RPC编辑 

使用 RPC 样式架构构建的基于 SOAP 的 Web 服务成为实现 SOA 最常用的方法。RPC 样式的 Web 服务客户端将一个装满数据的信封（包括方法和参数信息）通过 HTTP 发送到服务器。服务器打开信封并使用传入参数执行指定的方法。方法的结果打包到一个信封并作为响应发回客户端。客户端收到响应并打开信封。每个对象都有自己独特的方法以及仅公开一个 URI 的 RPC 样式 Web 服务，URI 表示单个端点。它忽略 HTTP 的大部分特性且仅支持 POST 方法。
由于轻量级以及通过 HTTP 直接传输数据的特性，Web 服务的 RESTful 方法已经成为最常见的替代方法。可以使用各种语言（比如 Java 程序、Perl、Ruby、Python、PHP 和 Javascript[包括 Ajax]）实现客户端。RESTful Web 服务通常可以通过自动客户端或代表用户的应用程序访问。但是，这种服务的简便性让用户能够与之直接交互，使用它们的 Web 浏览器构建一个 GET URL 并读取返回的内容。
在 REST 样式的 Web 服务中，每个资源都有一个地址。资源本身都是方法调用的目标，方法列表对所有资源都是一样的。这些方法都是标准方法，包括 HTTP GET、POST、PUT、DELETE，还可能包括 HEAD 和 OPTIONS。
在 RPC 样式的架构中，关注点在于方法，而在 REST 样式的架构中，关注点在于资源 —— 将使用标准方法检索并操作信息片段（使用表示的形式）。资源表示形式在表示形式中使用超链接互联。
Leonard Richardson 和 Sam Ruby 在他们的著作 RESTful Web Services 中引入了术语 REST-RPC 混合架构。REST-RPC 混合 Web 服务不使用信封包装方法、参数和数据，而是直接通过 HTTP 传输数据，这与 REST 样式的 Web 服务是类似的。但是它不使用标准的 HTTP 方法操作资源。它在 HTTP 请求的 URI 部分存储方法信息。好几个知名的 Web 服务，比如 Yahoo 的 Flickr API 和 Delicious API 都使用这种混合架构。 