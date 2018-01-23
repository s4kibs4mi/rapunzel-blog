![](https://raw.githubusercontent.com/s4kibs4mi/rapunzel-blog/master/extras/rapunzel_blog_thumb.png)
# Rapunzel Blog
A Blogging platform based on gRPC written in Golang
* status : development
* version : alpha-0.0.1

#### Features
`[+]` Registration<br/>
`[-]` Initiate Verification via Email<br/>
`[-]` User Verification via Email<br/>
`[+]` Login<br/>
`[-]` Get User Profile<br/>
`[-]` Change User Status<br/>
`[-]` Change User Type<br/>
`[-]` Change Password<br/>
`[-]` Reset Password Request<br/>
`[-]` Reset Password<br/>
`[-]` Logout<br/>
`[+]` Create Post<br/>
`[-]` Update Post<br/>
`[-]` Delete Post<br/>
`[+]` List Posts<br/>
`[+]` Get Post<br/>
`[-]` Change Post Status<br/>
`[+]` Favourite Post<br/>
`[+]` Create Comment<br/>
`[-]` Update Comment<br/>
`[-]` Delete Comment<br/>
`[+]` List Comments<br/>
`[+]` Get Comment<br/>
`[-]` Favourite Comment<br/>

##### Compile Protos,
```
protoc --proto_path=. --go_out=plugins=grpc:. ./protos/*.proto
```

#### License
Distributed under [Apache 2](https://github.com/s4kibs4mi/rapunzel-blog/blob/master/LICENSE) license
