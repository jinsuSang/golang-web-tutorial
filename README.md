# golang-web-tutorial
- Tucker의 Go 언어 프로그래밍 chapter 29 Go 언어로 만드는 웹 서버

## MUX

multiplex 약자로 여러 입력 중 하나를 선택해서 반환하는 디지털 장치이다. node.js 의 express 라이브러리의 router 와 유사합니다.

## static

- `http://localhost:3000/static/photo.jpg`

  ```go
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  ```

- static 폴더 지정 

  ```go
  http.FileServer(http.Dir("static"))
  ```

## Test Http

- testify 패키지 사용

1. response, request 생성 및 경로 테스트

   ```go
   res := httptest.NewRecorder()
   req := httptest.NewRequest("GET", "/", nil)
   ```

2. 핸들러 인스턴스 호출

   ```go
   mux := MakeWebHandler()
   mux.ServeHTTP(res, req)
   ```

3. code 확인 및 데이터 읽기

   ```go
   assert.Equal(http.StatusOK, res.Code)
   data, _ := io.ReadAll(res.Body)
   ```

4. 결과 확인

   ```go
   assert.Equal("Hello World!", string(data))
   ```

## JSON

1. Student 객체를 `[]byte`로 변환

   ```go
   data, _ := json.Marshal(student)
   ```

2. JSON 포맷 표시

   ```go
   writer.Header().Add("content-type", "application/json")
   ```

3. 전송

   ```go
   fmt.Fprint(writer, string(data))
   ```

   
