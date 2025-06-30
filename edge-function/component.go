package main

import (
	incominghandler "example-go-component/internal/wasi/http/incoming-handler"
	wasihttp "example-go-component/internal/wasi/http/types"

	"go.bytecodealliance.org/cm"
)

func Handle(request incominghandler.IncomingRequest, responseOut incominghandler.ResponseOutparam) {
	index := `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Coming Soon</title>
  <style>
    body {
      margin: 0;
      padding: 0;
      font-family: system-ui, sans-serif;
      display: flex;
      justify-content: center;
      align-items: center;
      text-align: center;
      height: 100vh;
      background-color: #f4f4f4;
      color: #333;
    }
    .container {
      max-width: 400px;
      padding: 2rem;
    }
    h1 {
      font-size: 2.5rem;
      margin-bottom: 1rem;
    }
    p {
      font-size: 1.1rem;
      margin-bottom: 2rem;
    }
    footer {
      font-size: 0.9rem;
      color: #888;
    }
    a {
      color: #007bff;
      text-decoration: none;
    }
    a:hover {
      text-decoration: underline;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Coming Soon</h1>
    <p>We're working hard to launch something awesome. Stay tuned!</p>
    <footer>Served by <a href="https://www.edgee.cloud">Edgee</a></footer>
  </div>
</body>
</html>
`
	bytes := []uint8(index)
	/*
	   let resp_tx = OutgoingResponse::new(self.headers);
	   let _ = resp_tx.set_status_code(self.status_code);

	   let body = resp_tx.body().unwrap();
	   ResponseOutparam::set(resp, Ok(resp_tx));
	   let stream = body.write().unwrap();
	   if let Some(body_content) = self.body_content {
	       let _ = stream.write(body_content.as_bytes());
	   }
	   drop(stream);
	   let _ = OutgoingBody::finish(body, None);
	*/
	response := wasihttp.NewOutgoingResponse(wasihttp.NewFields())
	response.SetStatusCode(200)
	body, _, _ := response.Body().Result()
	stream, _, _ := body.Write().Result()

	wasihttp.ResponseOutparamSet(responseOut, cm.OK[cm.Result[wasihttp.ErrorCodeShape, wasihttp.OutgoingResponse, wasihttp.ErrorCode]](response))

	index2 := cm.NewList(&bytes[0], len(bytes))
	stream.Write(index2)
	stream.BlockingFlush()
	stream.ResourceDrop()
	_, _, iserr := wasihttp.OutgoingBodyFinish(body, cm.None[wasihttp.Trailers]()).Result()
	if iserr {
		panic("Failed to finish outgoing body")
	}
	println("Response finished successfully")
}
