# fire

更好用、能扩展、支持多国语言提示的表单验证类库，使用`GoLang`编写。

## 用法

我们来看一个最简单的例子，这是`fire`类库的`Hello World`。

```golang
import (
    "fmt"
    "github.com/joyant/fire"
)

func main() {
    v := fire.New(fire.Rule{
        "name":"required|lengthBetween:5,10",
    })
    data := fire.Data{"name":"Tom"}
    qulified, err := v.Validate(data)
}
```
我们先有一个直观的感觉：使用`fire`是简单的，下面是更具体的例子。

### 传入fire.Data

`fire.Data`是`fire`定义的数据类型，专门用于传入数据用，我们可以直接使用它。

```golang
v := fire.New(fire.Rule{
    "name":"required",
})
data := fire.Data{"name":"Tom"}
v.Validate(data)
```
`fire.Data`本质上一个`map[string]interface{}`类型，所以我们也可以传入`map[string]interface{}`类型。

```golang
v := fire.New(fire.Rule{
    "name":"required",
})
data := map[string]interface{}{"name":"Tom"}
v.Validate(data)
```

### 传入struct
```golang
type User struct {
    Name string
}
v := fire.New(fire.Rule{
    "name":"required",
})
u := User{Name:"Tom"}
v.Validate(u)
```
传入指针也是可以的，但是指针必须指向一个结构体。
```golang
u := &User{Name:"Tom"}
v.Validate(u)
```

需要注意的是，除了`fire.Data`、`map[string]interface{}`、结构体以及指向结构体的指针外，`Validate`方法并不支持传入其他类型的参数，如果传入了错误类型的参数，返回的`err`不为空，检查`New`函数返回的`err`是否为空是有必要的。

### 结构体(struct)与Rule的映射关系

在上一部分的例子里，我们看到结构体的属性是`Name`,而`Rule`的`key`是`name`，严格来说，它们并不相等(用==判断)，那么`fire`是如何将他们对应起来的呢？

`fire`会先将`struct`的属性按照一定的规则转为`Rule`的`key`，当这些规则都用完了还没有找到映射关系，那`fire`就会
认为这个数据并不存在，规则如下：

```golang
Name       -> Name
Name       -> name
FirstName  -> first_name
First_Name -> first_name
```
当`fire`拿到一个`key`，它会依次寻找其原始、全小写、下划线形式在`Rule`是否有对应的`key`，如果都没有找到，就会认为其对应的规则是空，也就是说这个属性是不会被校验的。

假设传入`Validate`函数的参数是`fire.Data{"name":"Tom"}`,规则是`Rule{"Name":"required","name":"int"}`,`fire`会依次在`Rule`中寻找`Name`、`name`，先找到谁，就认定那个验证规则，在这个例子里，规则就是`required`,而不是`int`。

在`GoLang`的世界里，很多验证类库的方式都利用了`tag`，这样做的好处是，`key`的定义非常清晰，没有歧义；但也有弊端，它会使结构体中包含很多让人眼花缭乱的`tag`，比如：

```golang
type User struct {
    Name string `json:"name", validator:"name", form:"name", db:"name"`
}
```
在上面的结构体中，`json`标签是用来转字符串或反转用的，`validator`标签是验证表单用的，`db`和`form`也各有其用途，本来是一个"纯粹"的结构体，在加入很多的`tag`后，结构体变得非常臃肿；我们在`fire`中偏爱"约定优于配置"的原则，这些约定都是比较符合一般做法和直觉的，不会令人觉得怪异，这样就不大需要写`tag`了。

但是，我们的约定可能无法满足所有的要求，因为我们不总是从一个的项目的零步开始做起，很可能接手的是一个老项目，而那里大概率已经有一些约定俗称的规则了，我们总不能因为引入了一个表单验证的类库就打破原来的规则，所以，`fire`也是支持`tag`的，当约定不能满足要求的时候，可以将`tag`应用于特殊的需求，`fire`给了`tag`最高的优先级，请看下面的例子：

```golang
type User struct {
    Name string `fire:"nickname"`
}
```
这样`name`对应的`key`就是`nickname`，即使`data`有另外一个名为`name`的`key`，`fire`也不会去验证它，因为我们
给了`tag`最高的优先级。

如果你不喜欢`fire`这个默认的`tag`（也可能是想复用名为`json`的`tag`），可以通过以下设置修改：
```golang
fire.Tag = "json" // db ...
```



## 验证规则

我们把内置的规则都列出来：

token | 用途 | 用法举例 | 意义 | 备注
-|-|-|-|-
alphaNum | 英文字符或数字 | alphaNum | 英文字符或数字 | 
alpha | 英文字符 | alpha | 英文字符 |
between | 数字大小必须在指定范围内 | between:1, 10 | 数字必须在1~10之间 |1和10也可以是浮点型
bool | 必须是布尔值 | bool | 必须是布尔值 |true或false的字符串表示
contains | 必须包含指定字符 | contains:abc、contains:abc/i | 字符串里必须包含abc |abc连在一起时，才能通过验证，abc/i表示不区分大小写，含有ABC的字符串也能通过验证
date | 必须是日期格式 | date/date:2006-01-02/date:2006-01-02 15:04:05等 | 必须是date后面指定的日期格式，不指定格式默认是2006-01-02 15:04:05格式 | 2006、2006-01、2006-01-02、2006-01-02 15:04:05都可以作为date后面的格式 
different | 必须和指定key对应的值不同 | different:username | 字符串值不能和key为username的值相同|常用于如判断密码不能和账号相同
email | 邮箱格式 | email | 邮箱格式 | 
equals | 必须与指定key的值相同 | equals:confirm_password | 必须与指定key的值相同 |常用于二次输入密码和第一次输入是否相同的验证
in | 在指定字符串之内 | in:1,2,3 | 必须在[1,2,3]之内|会自动trim逗号之间的空白
notIn | 不能在指定字符串之内 | in:1,2,3 | 不能在[1,2,3]之内 |会自动trim逗号之间的空白
integer/int | 必须是整型 | integer或者int | 必须是整型 |
ip | 必须是ipv4或ipv6格式 | ip | 必须是ipv4或ipv6格式 |
ipv4 | 必须是ipv4格式 | ipv4 | 必须是ipv4格式 |
Ipv6 | 必须是ipv6格式 | Ipv6 | 必须是ipv6格式 |
lengthBetween | 字符串长度在指定范围之内 | lengthBetween:1,10 | 字符串长度在1~10之间 |utf8编码
lengthMax | 字符串长度最大范围 | lengthMax:10 | 字符串长度不能超过10 |utf8编码
lengthMin | 字符串长度最小范围 | lengthMin:10 | 字符串长度不能小于10 |utf8编码
length | 字符串长度必须是指定值 | length:5 | 字符串长度必须是5 |utf8编码
max | 最大值 | max:100 | 值不能超过100 |内部实现是转成float64再比较的
min | 最小值 | min:10 | 值不能小于10 |内部实现是转成float64再比较的
numeric | 必须是数字（整数或小数） | numeric | 必须是数字（整数或小数） |100和99.99都能通过验证
regexp | 必须匹配指定正则表达式 | regexp:\\\d+{1,10} | 匹配1~10个数字 |
required/require | 必填项 | required | 必填项 |
url | 必须是网址格式 | url | 必须是网址格式 |http://www.abc.com，https://www.abc.com，http://abc.com,https://abc.com,www.abc.com,abc.com都可以通过验证，后面跟路径也可以通过验证

需要注意的是，规则之间要用`|`分开，并且冒号要用英文，不能用中文，比如我们可以设定以下规则来验证：

```golang
rule := Rule{
    "username":"in:Tom,Jim,Jerry|lengthBetween:1,10|email|int|numeric|alpha",
}
```

这只是个例子，只是想让大家知道多个验证规则是被支持的，只要用`|`作分隔符就可以了，事实上，例子里的规则不能用来校验任何东西。



## 内置的中英文提示

不能验证通过时返回的提示信息非常重要，它能明确告诉用户哪个字段不能同通过验证，而不是笼统告诉用户发生错误了。在`fire`内置的这些规则中，都有中英文提示，比如我们要验证一个`key`为`birthday`的值，它不能通过日期格式`(date)`验证时，返回的中文提示是`date必须是日期格式`,英文提示是`date must be date format`。在某些时候，我们可能不希望用户看到中英文掺杂的文字提示，换句话说，我们想让`birthday`有指定的中文翻译，这时，我们就可以调用`fire`的`RegisterI18nDataKey`函数：

```golang
fire.RegisterI18nDataKey(fire.DataKey("birthday"), map[fire.Lang]string{
    fire.LangZH:"生日"
})
// 也可以使用fire.RegisterI18nDataKeyMap函数来同时注册多个多语言支持
```

这个函数的作用是为`birthday`注册多语言的翻译，这个时候我们就能看到全中文提示了`生日必须是日期格式`,如果你的用户本来就是英文世界的，那就不需要调用以上的函数了，只需要把默认提示语言设置成英文就行了，设置方法有两种：

```golang
//第一种方法,是全局设置
fire.DefaultLang = fire.LangEN 
//第二种方法是以validator为单位设置，这个validator返回的验证结果将会是英文提示，不会影响全局设置
fire.New(fire.Rule{}, fire.LangEN) 
```

`fire`的默认语言设置是中文，即使正合你的需要，还是建议设置默认语言。我们内置了一些常量供开发者使用，这些常量是为开发扩展和设置消息提供方便，并不是说`fire`内置了这些语言的提示，内置的只有中文和英文。

```
LangZH Lang = "zh" // 汉语
LangEN Lang = "en" // 英语
LangDE Lang = "de" // 德语
LangFR Lang = "fr" // 法语
LangRU Lang = "ru" // 俄语
LangES Lang = "es" // 西班牙语
LangJA Lang = "ja" // 日语
LangAR Lang = "ar" // 阿拉伯语
LangKR Lang = "kr" // 韩语
LangPT Lang = "pt" // 葡萄牙语
```



## 多国语言支持

假设你的项目有不少中东用户，所以需要一套阿拉伯文的验证结果提示，但是内置的语言只有中文和英文，该如果解决这个问题呢？我们还拿生日来举例子，直接看代码：

```golang
fire.RegisterMsgFormat("birthday", map[fir.Lang]fire.MsgFormat{
    fire.LangAR:"${0} يجب أن يكون شكل عيد ميلاد ", //birthday必须是生日格式
})
```

这是我们第一次接触到信息提示的占位符，在以上的例子里`${0}`就代表`birthday`，关于占位符，我们会在扩展验证部分更多提到。



## 扩展验证规则

虽然`fire`内置了常用的验证规则，但一定会遇到不够用的时候，在这种情况下，`fire`允许用户扩展验证规则，假设我们现在有一个比较奇怪的规则：验证姓名，如果是英文名字，那必须得是全名(full name)，其他名字（中文，阿拉伯文等）必须满足指定的最小长度才可以通过验证；这个规则显然是常用的规则不能满足要求的，那我们来扩展一个规则吧。

### 原理

希望所有的扩展开发者都能知道扩展的原理是什么，这样一定写得更好。在`fire`内部，每个验证规则都对应了一个`Token`实例，`Token`是一个接口，有一系列规定的方法，`fire`在验证数据的时候，就是调用`Token`的一连串方法来判断的，所以我们只需要注册一个新的`Token`让`fire`知道就可以了。

### 代码

```golang
type specialNameToken struct {
	literal string
    length int
}

func (t *specialNameToken) Evaluate(value DataValue, data Data) 
(qualified bool, literalValue []string, err error) {
    if value == nil { //如果传入的数据是空，那我们就不验证了
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    isENName, includeSpace := isName(v)
    if isENName {
        return includeSpace, []string{t.literal}, nil
    } else {
        return utf8.RuneCountInString(v) >= t.length, []string{t.literal}, nil
    }
}

func (t *specialNameToken) TokenType() TokenType {
    return "specialName" //token独一无二的名称,不应该和其他规则重复
}

func (t *specialNameToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是指定格式" // ${0}是一个占位符，将会被替换成数据的key
    } else if lg == LangEN {
        return "${0} must be special format"
    }
    return ""
}
//ParseLiteral接收到的参数是specialName:后面的值
func (t *specialNameToken) ParseLiteral(literal string) error {
	if literal != "" {
        length, err := strconv.Atoi(literal)
        if err != nil {
            return err
        }
        t.length = length
	}
	t.literal = literal
    return nil
}

//clone用来保证深拷贝一个对象，修改拷贝时不会影响原来的对象
func (t *specialNameToken) Clone() Token {
    c := *t
    return &c
}

func isName(s string) (isENName bool, includeSpace bool) {
    for _, v := range s {
        if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || v == ' ') {
            return false, false
        }
        if v == ' ' {
           includeSpace = true
        }
    }
    isENName = true
    return
}
```

我们完成了一个扩展，并且规定了其名称是`specialName`，名称是`TokenType()`的返回值决定的；但这个`token`是`fire`所不认识的，我们把它注册给`fire`，然后就可以放心的使用了：

```golang
fire.RegisterToken(&specialNameToken)

rule := fire.Rule{
    "name":"specialName:2",
}
```

### 占位符

按我们约定的，现在详细的来看下占位符，我们拿`fire`内置的一个验证规则来举例子,比较有代表性的规则是`lengthBetween`:

```golang
// fire.Rule{"name":"lengthBetween:1, 10"}
func (t *lengthBetweenToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}长度必须在${1}和${2}之间"
    } else if lg == LangEN {
        return "${0}'s length must between ${1} and ${2}"
    }
    return ""
}
```

`${0}`将会被`name`填充，`${1}`将会被`1`填充，`${2}`将会被`10`填充。

我们为什么要解释占位符呢？因为占位符不但在开发扩展的时候有用，在设置多国语言提示的也是有用的。在"多国语言"部分，我们提到了如果我们的项目是中文和英文世界之外的用户使用，我们如何给他们提示相应的方言呢？看代码：

```golang
fire.RegisterMsgFormat("birthday", map[fir.Lang]fire.MsgFormat{
    fire.LangAR:"${0} يجب أن يكون شكل عيد ميلاد ", //翻译过来就是：birthday必须是生日格式
})
```

这样我们就注册了一个`date`规则对应的消息提示，当错误发生时，就会返回阿拉伯文的提示，当然，前提是设置了默认的语言。

## 别名

请看以下代码:

```golang
fire.RegisterI18nDataKey(fire.DataKey("class_name"), map[fire.Lang]string{
    fire.LangZH:"班级名称"
})
v := fire.New(fire.Rule{
    "name":"alias:class_name|required"
})
```

由于`name`这个关键词实在太宽泛了，它在`user`表中表示用户名，在`class`表中表示班级名称，在`teacher`表中又表示教室名称，但是在不同的接口中，它们都有一个相同的名称：`name`，我们可不希望每个接口都提示为`名称`，因为这个提示太不具体了，我们还是希望提示："名称不能为空","班级名称不能为空","教室名称不能为空"……，所以`fire`内置了名为`alias`的规则，在上面的例子里，当`name`不能通过验证时，提示的是"班级名称是必填项"。

## 最佳实践

在`fire`的设计里，我们故意把解析`Rule`和验证数据分开了，当调用`fire.New`方法返回一个接口时，已经把规则都解析好，"缓存"起来了，剩下的就是数据的验证了，因为验证本身不会修改对象的成员变量，所以同一个对象的`Validate`方法可以在多个协程中同时调用，而不会因为并发产生问题。所以比较好的实践是不要每次收到请求时都去调用`fire.New`,而应该在项目初始的时候就把对象实例化好，在需要验证数据的时候，调用指定对象的`Validate`方法就可以了，这减少了很多解析规则的开销，举个简单而正确的例子，假设以`handleAPI`为前缀的函数运行在不同的协程中：

```golang
var v1 = fire.New(rule1)
var v2 = fire.New(rule2)

func handleAPI1(data fire.Data) {
    v1.Validate(data)
}
func handleAPI2(data fire.Data) {
    v2.Validate(data)
}
```

下面是错误的例子：

```
func handleAPI1(data fire.Data) {
	var v1 = fire.New(rule1)
    v1.Validate(data)
}
func handleAPI2(data fire.Data) {
	var v2 = fire.New(rule2)
    v2.Validate(data)
}
```

在每个请求到来时都创建`fire.Validator`对象是没有必要的，增加了不必要的开销。

在`fire`里，那些`Register`开头的函数，实际的作用是把一些我们需要的数据以键值对的形式存放在`map`里，`fire`使用的就是`GoLang`内置的`map`，没有加锁，不支持同时读写，因为注册类的函数调用应该在调用`fire.New`前进行，一旦有`fire.Validator`实例被创建出来了，任何注册类函数都不应该再被调用了。



## 起源和思路

在我看到一个开源项目时，尤其是要在项目中使用时，总是非常想和项目的作者谈一谈，希望了解项目的设计思路、原因和目的等，想知道这些项目背后的起源，我觉得那是历史的一部分，这其中的故事就像其他伟大的历史事件一样吸引我，所以我想说一点额外的东西，如果你只是想知道`fire`如何使用，我已经说完了，到这就可以停下了。

表单验证类库在`web`应用中是如此重要，它是业务逻辑的起始，我自己在开发项目时，总倾向于用一种尽量简单的方法验证数据的正确性，而把更多的精力用在真正业务逻辑的编写上。受到`thinkphp`框架里验证方法的启发，就想把简单的方法带到经手的项目中来，这就是`fire`的来源。

在开发`fire`的过程中，我也从解析器中借鉴了一些方法，把验证规则变成可验证的函数，其实就是把字符串解析成有特定目的的数据结构。所以先确定下了几个重要接口：`Token`,`parser`,`Validator`，明确了他们各自的职责，`Token`负责制定规则并验证数据，`parser`被`Token`调用把字符串转为有意义的数据结构，就像解释有指定语法的代码，`Validator`负责把数据传给相应的`Token`，让它去验证，最后把结果返回给调用者。

得益于`GoLang`的接口，先定义接口，写出主要的业务逻辑，写单元测试，最重要的工作并没有花费很多的时间。剩下的工作就是完成一个个内置`Token`的开发，开发完后注册到保存`Token`的`map`里，测试。

所以，`fire`最重要的并不是那些内置的规则，而是基于接口的一系列调用，即使`fire`中一个内置的规则都没有，开发者还是可以通过自己开发`Token`然后注册到`fire`来使用，这些`Token`甚至可以覆盖内置的`Token`，开发者非常可能比我写得更好。