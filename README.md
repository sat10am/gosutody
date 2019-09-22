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
<summary>190922-9th</summary>
  
- Concurrency
  - 여러개의 웹사이트를 비동기적으로 체크하는 CheckWebsites를 만들면서 고루틴의 기본을 배움
  - 고루틴은 `go`로 선언하고 Go에서 동시성을 다루는 기본 단위이다.
  - 각각의 웹사이트를 체크하는 프로세스를 고루틴으로 실행하기 위해 익명함수를 사용했다.
  - 테스트를 PASS했어도 여러번 돌려보면 랜덤하게 PANIC 에러가 뜨는데, 경합(race)이 발생하기 때문
  - Channel이라는 자료구조를 사용하여 경합을 해결할 수 있다.
  - race 옵션으로 동시성 코드에서 발생할 수 있는 문제를 디버깅 가능하다.
  - Make it work, make it right, make it fast

</details>

<details>
<summary>190915-8th</summary>
  
- Mocking
  - countdown 프로그램을 만들고 timer를 Mocking했다.
  - 문제를 작게 분할하고 동작하는 프로그램을 먼저 만들어라.
  - 코드가 어려워서 테스트하기 어렵다고? 테스트하기 좋게 코드를 짜야 한다.
  - 리팩토링은 코드는 달라도 동작은 같은 것이다.
  - 테스트의 가치와 리팩토링했을 때의 효과를 항상 염두에 둬야 한다.
</details>

<details>
<summary>190908-7th</summary>
  
- Dependency Injection
  - DI 를 사용하면, 테스트 친화적이고 범용적인 함수를 작성할 수 있다
  - 만약 함수의 테스트가 쉽지 않다면, 전역 상태나 의존성에 대해 강하게 묶여 있다는 의미이다
    - 전역 데이터페이스 연결 풀 같은 경우에는 테스트 하기가 쉽지 않은데 (서비스 레이어에서 사용되고 있으므로) 이런 경우에 활용할 수 있다
  - 데이터가 생기고 흘러가는 것에 대한 디커플링이 가능합니다
    - 만약, 함수의 의존성이 너무 많다고 느껴진다면 DI가 올바른 답일 수 있습니다
  - 컨텍스트에 한정되지 않는 범용적인 함수를 작성할 수 있습니다
  - Go Built-In library 를 잘 살펴보는 것이 올바른 함수를 작성하는데에 도움이 될 

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
