namespace go api

struct EchoReq {
    1: string Message (api.query="message"); // 添加 api 注解为方便进行参数绑定
}

struct EchoResp {
    1: string Message;
}

service Echo {
    EchoResp EchoMethod(1: EchoReq req) (api.get="/echo")
}