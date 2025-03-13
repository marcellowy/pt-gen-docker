import {gen_douban} from "./lib/douban.js";
import {makeJsonResponse} from "./lib/common.js";

// curl -H "Content-Type: application/json;charset=utf-8" -d "{\"sid\":34429795}" http://localhost:3000/api/v1/ptgen
let response_data = makeJsonResponse();
try {
    let sid = process.argv[2];
    if (sid) {
        response_data = await gen_douban(sid);
    }
    if (response_data) {
        let response = makeJsonResponse(response_data)
        let body = await response.json()
        // 接口兼容处理
        body.ret = 200
        body.msg = "success"
        console.log(JSON.stringify(body));

    }
} catch (e) {
    console.log(e)
}
