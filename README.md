# golang-web-tutorial
Tucker의 Go 언어 프로그래밍 chapter 29 Go 언어로 만드는 웹 서버

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

  

