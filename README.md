## What is this study for?
일요일 아침 ~~11시에~~10시에 Go Study 를 하는 GOSU들의 저장소 입니다

## How to study
1. <[learn-go-with-tests](https://github.com/quii/learn-go-with-tests)> 문서의 Go fundamentals 예제를 따라 TDD로 Go를 실습하면서 문법을 익히고, 결과물을 자신의 디렉토리에 커밋합니다.
2. 책 <[디스커버리 Go 언어](http://www.hanbit.co.kr/store/books/look.php?p_code=B5279497767)> 는 필요한 경우 참조합니다.
3. 문법 공부가 끝나고 간단한 토이 프로젝트를 만들어 봅니다.

## History
<details>
<summary>sample</summary>
  
- nil

</details>

<details>
<summary>190901-6th</summary>
  
- TDD: 8.Maps
  - key-value 형태로 값을 저장할 수 있으며 key를 통해 빠르게 데이터 검색이 가능한 해쉬테이블
  - value 타입은 어떤 타입이든 들어올 수 있지만 key의 타입은 Comparable 타입만 가능하다.
  - 결과 값을 반환할 때 Error객체에 문자열을 담아 반환 할 수 있다.
  - map은 reference 타입이기 때문에 아무리 큰 맵이라도 한번에 복사가 가능하다.
  - map은 절대로 empty하게 초기화 하면 안된다. (make를 사용하거나 {}를 붙여 선언해야한다.)
  - Error는 불변성과 재사용성을 위해 상수로 선언해서 사용하자.
  - map의 delete는 내장 함수를 사용한다. 

</details>

<details>
<summary>190822-5th</summary>
  
- TDD: 7.Pointers & errors
  - call by value vs call by reference: 상태를 바꾸고 싶다면 pointer 사용
  - 미리 다양한 Error 타입을 만들어 두고 사용할 수 있다
  - Error를 체크만 하지 말고 처리하자
  - 포인터만 nil을 가질 수 있으며, C처럼 이중 포인터 사용 가능(속도 최적화시)
  - `_`([blank identifier](https://golang.org/doc/effective_go.html#blank))는 관습적으로 선언은 하지만 사용하지 않을 변수를 말한다. lint에서도 미사용에 대해 잡지 않음
  
</details>

<details>
<summary>190821-4th</summary>
  
- TDD: 6. Structs, methods & interfaces (2/2)
  - byte to uint64 어쩐지 타입캐스팅 (갑분C)
  - go polymorphism, encapsulation, pointer 조금씩 다룸
  
</details>

<details>
<summary>190817-3th</summary>
  
- TDD: 6. Structs, methods & interfaces (1/2)
  - go struct, methods
  - go 에는 (명시적인) 메소드가 없으며, struct 는 여러 형태의 필드의 묶음일 뿐이다
  - go struct 의 생성 방법은 여러가지가 있다. `new`, `literal struct` ...
  - go interface 를 진행하려가다가 공부를 더 하고 진행하는 방향으로 후퇴
- 다음 시간까지 숙제: struct, methods, interfaces 각자 정리해서 push (++ codekata 할 수 있으면 추가)

</details>

<details>
<summary>190815-2th</summary>
  
- TDD: 3. Integers / 4. Iteration / 5. Arrays and slices
  - go Benchmark, Test coverage
  - go doc -> 아직 모듈에서는 제공하지 않음
  - reflect.DeepEqual 제공 (but not type-safety)
- 다음 시간까지 숙제: Min, Max, Sort(bubble부터 quick까지 아무거나) 함수 만들기

</details>

<details>
<summary>190811-1th</summary>
  
- TDD: 0. Install / 1. Hello, world
  - GVM으로 Go 설치 (먼저 1.4 설치 후 1.12.7버전 사용)
  - [gvm이용하여 go설치하기](https://select995.netlify.com/go/module/gvm)
- 다음 시간까지 숙제: TDD로 간단한 사칙연산 계산기 만들기

</details>

## References
- [learn-go-with-tests](https://github.com/quii/learn-go-with-tests)
- [gvm](https://github.com/moovweb/gvm)
- [TDD로 Golang 시작하기 by y0c](https://y0c.github.io/2019/08/11/beginning-go/)
- [디스커버리 Go 언어](http://www.hanbit.co.kr/store/books/look.php?p_code=B5279497767)
