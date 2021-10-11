# golang-web-tutorial
- Tucker의 Go 언어 프로그래밍 chapter 29 Go 언어로 만드는 웹 서버

- Tucker의 Go 언어 프로그래밍 chapter 30 RESTful API 서버 만들기
- Tucker의 Go 언어 프로그래밍 chapter 31 Todo 리스트 웹 사이트 만들기
- https://todolist-ruslan-lvivsky.herokuapp.com/

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

## RESTful API 특징

- 자기 표현적인 URL

  URL 만으로 어떤 데이터를 요청하는지 인식

- 메서드로 행위 표현

  메서드와 URL로 데이터 조작을 정의

- 서버 클라이언트 구조

- stateless

  서버가 상태를 보관하지 않음

- cacheable 캐시 처리

  더 쉽게 캐시 정책을 적용하여 성능을 개선

## RESTful Update

- GET 형식으로 업데이트 양식 받기 => 보안에서 위험함

- 모든 값을 string 형식으로 받기 => 데이터 형식이 두가지가 형성됨 

- `map[string]interface{}` 형식으로 하고 나중에 Type Assertion => 

  반복문 형식으로 key 로 하고 value.(int) 로 하려 하였으나 타입이 한 번밖에 타입 단언되지 않음

- update 관련된 구조체를 사용하여 업데이트 여부 확인 => 구조체 추가 및 불필요한 데이터 증가

  

## Todo List

### urfave/negroni 패키지

- 로그 기능 

  웹 요청을 받아 응답할 때 자동으로 로그를 남겨 웹 서버 동작을 확인한다

-  panic 복구 기능

  웹 요청을 수행하다가 panic이 발생하면 자동으로 복구하는 동작을 지원한다

- 파일 서버 기능

  public 폴더의 파일 서버를 자동으로 지원한다

### unrolled/render 패키지

- 웹 서버 응답을 구현하는데 사용하는 유용한 패키지이다

### 클라우드 서비스 유형

#### IaaS 서비스

Infrastructure as a Service -  CPU, 메모리, 저장 장치, 네트워크 같은 웹 서비스에 필요한 하드웨어 장비를 제공한다. 

#### PaaS 서비스

Platform as a Service - 서비스가 플랫폼 형태로 제공한다. 사용량에 따라 성능을 맞추고 웹 서버 관리 및 실행을 담당한다 

#### SaaS 서비스

Service as a Service - 웹 서비스를 통해서 소프트웨어를 사용하는 형태이다. 정해진 포맷으로 웹 사이트를 빠르게 지원하는 서비스도 이에 속한다
