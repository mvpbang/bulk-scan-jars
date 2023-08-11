# 需求：
- 01、读取目录下war、ear、jar，获取jar路径及格式化出jar_name + jar_version
- 02、和安全jar名字及坐标做比较，判断出不安全的
- 03、输出以csv形式(,分割字段)
- 04、仅仅对需要搜的jar做名字和版本提取

## 支持正则匹配的版本
```
//jar信息
axis2-metadata-1.7.8.jar
BondeReverseDesensitize-0.0.25-SNAPSHOT.jar
nacos-spring-context-0.0.2.jar
nacos-spring-context-0.0.2-20230718-SNAPSHOT.jar
x-spring-context1-0.0.2-20230718-SNAPSHOT.jar
commons-collections4-0.0.2.jar
javax.inject-1.jar
geronimo-annotation_1.0_spec-1.1.jar 
geronimo-jaxws_2.2_spec-1.0.jar
serializer-j_2_7_2.jar

//正则

#v0.0.0
^(?P<name>.*?)-\d
\b(?:\d+\.?)+\-?\d+|\d

#v0.0.1
^(?P<name>.*?)-\d
-\b(?:\d+\.?)+\-?\d+|-\d  //-作为界定符,对空值做特殊处理

共找到 9 处匹配：
-1.7.8
-0.0.25
-0.0.2
-0.0.2-20230718
-0.0.2-20230718
-0.0.2
-1
-1.1
-1.0
```