# PI DASH BackEnd
Pi Dash 프로젝트를 위한 데스크톱 백엔드

## 개요
데스크톱의 성능 지표를 간단히 확인할 수 있는 PiDash 프로젝트의 백엔드\
실행되는 장치의 성능 지표들을 HTTP/WS를 통해 제공, 이 정보를 PiDash APP에서 수신해서 보여줌

## 개발환경
- Go
    - gorilla
        - mux
        - websocket
    - gopsutil
    - net/http

## 라우트 정보
- /api
    - /cpu : CPU 정보 (vendorId, family, model, clock 등) 제공
    - /gpu : GPU 정보 (ProductName, DriverVersion, TotalMemory, MaxClocks) 제공
    - /mem : Memory 정보 (전체 용량) 제공
    - /net : Network 정보 (인터페이스 명, Mac Addr) 제공
- /ws
    - /cpu : 실시간 CPU 정보 (점유율) 제공
    - /gpu : 실시간 GPU 정보 (GPU/Mem 사용량, 온도, 전력, Clock) 제공
    - /mem : 실시간 Memory 정보 (Total, Available, Used, Free) 제공
    - /net : 실시간 Network 정보 (인터페이스 명, Tx/Rx 속도) 제공

## ToDo
1. CPU Temp 정보 수신
2. Windows 외 타 OS 지원 추가