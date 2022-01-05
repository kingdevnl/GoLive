<div>
    <h1>This is a test.</h1>
    <p>{{.message}}</p>



    <p>c:{{.counter}}</p>

    <button live-click="clicky">ClickMe!</button>


    <p>search: {{.search}}</p>
    <input type="text" placeholder="search" live-input="search" value="{{.search}}">



{{/*    {{live_child "ChildComponent"  "test1" .state }}*/}}
</div>