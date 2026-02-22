
게임 개발의 업계 스탠더드는 "게임 엔진"을 사용한 게임 개발입니다.
게임 엔진의 대표로는 아래 친구들이 있습니다.   
- [Unreal](https://www.unrealengine.com/ko/unreal-engine-5?gad_source=1&gad_campaignid=9593971561&gclid=CjwKCAiAzOXMBhASEiwAe14SaUePc9vOLY3_HGOcDv055801lRek0OTKJy4_bDv-2ioHlweUNo-76BoC4u4QAvD_BwE)
- [Unity](https://unity.com/kr/products/unity-industry?utm_source=google&utm_medium=cpc&utm_campaign=nxt_abm_ind_apac_kr_ko_co_sem-gg_acq_br-pr_2026-01_unityindustry-evergreen_cc3022_x&utm_term=unity&gclsrc=aw.ds&gad_source=1&gad_campaignid=23425396452&gclid=CjwKCAiAzOXMBhASEiwAe14SaSgewnQmaeFveSpjNlNhkVoP6qEZDBK7YUeNu0N_Yr9X58_AmxF07RoCbQUQAvD_BwE)
- [Godot](https://godotengine.org/)

게임 엔진이 발전한 이유는, 게임 개발이라는 행위에 필요한 기술이 많아서 그렇습니다.
- 그래픽 API 
- 음악 창작 
- 에셋 창작 
- 프로그래밍  

하나하나 만만한 영역이 없기에, 개인이 시도하기엔 무리가 있고, 
그 무리가 있는 부분을 게임 엔진이 담당해줍니다. 

이중 가장 hard한 예시 하나만 들어보겠습니다.  

그래픽스 API 호환성 이슈:  
그래픽스 API는 파편화 되어 압도적인 표준이 없습니다.  
오늘날 그래픽스 API는 하나의 표준이 지배하는 대신, 하드웨어와 OS에 최적화된 '저수준(Low-level) API'들로 쪼개져 있습니다.  
- Windows / Xbox:	DirectX	윈도우 게임 개발 표준    
- Android / Linux:	Vulkan	OpenGL의 후계자  
- Apple (Mac/iOS):	Metal	애플 생태계 전용  
- Web:	WebGPU	WebGL(OpenGL 기반)을 대체하기 위해 등장한 차세대 웹 그래픽 표준  

개인이 이부분을 해결하려면, 게임 개발 하기 전에 이 표준들부터 전부 공부 해야합니다.
공부해서 내 게임에 호환성을 부여해야, 각 플랫폼에 배포할 수 있습니다.
참고로 openGL을 대체 하기 위해 나온 Vulkan은 하드웨어적인 세부적인 제어가 가능하기 때문에,
코드가 훨씬 방대합니다. 
OpenGL은 '그려라'라고 명령하면 드라이버가 알아서 처리해주지만, Vulkan은 '메모리를 할당하고, 동기화하고, 하드웨어 상태를 직접 관리하며 그려라'라고 해야 합니다.
예시를 하나 붙여두겠습니다.  
[vulkan_tutorial_github](https://github.com/SaschaWillems/HowToVulkan)

그래서 현실적으로 그래픽 프로그래머가 인생 목표가 아닌 사람이 상업적인 게임 개발이 목표라면  
게임 엔진을 사용하게 됩니다.  

근데, 게임 엔진도 문제가 많습니다.
- 숨겨진 동작 
- 불편한 UI
- 버전 업데이트 때마다 대격변 
- 특정 게임 엔진을 사용해 게임을 시장에 내면, 수수료를 내야 함 (Unreal, Unity)

만약 게임 개발을 시작하는 사람이 프로그래머라면, 택할 수 있는 제 3지대가 있습니다.  
- 어렵고, 게임 개발과는 좀 거리가 있는 부분만 처리 해줌  
- 기초적인 기능을 API 형태로만 제공하고, 내가 모두 개발  
프레임워크, 라이브러리 등 부르는 사람 마음대로인데 .. 예시 기술을 공유 드리겠습니다.  
- [SDL](https://www.libsdl.org/)
- [libGDX](https://libgdx.com/)  
- [raylib](https://www.raylib.com/): *cheat sheet를 적극 활용하시면 좋습니다.

저는 raylib을 선택했으며, 이 예제를 통해 배움을 공유하려고 합니다.  
raylib의 go wrapper 코드를 가져와서 사용합니다.  
[raylib-go](https://github.com/gen2brain/raylib-go)
- C는 현대에 사용하기엔 좀 불편하고, Zig는 아직 성숙하지 못하다고 판단해서 제외 했습니다.  
- Go는 배워두면 직장의 기회도 얻을 수 있다고 판단하여 Go를 선택 했습니다.

시작 하기: 
1. 사용하는 운영체제에 따라 요구사항 설치
2. go get -v -u github.com/gen2brain/raylib-go/raylib 실행
3. main에 example 코드 작성
