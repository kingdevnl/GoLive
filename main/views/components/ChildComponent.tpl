<div>
    <h1>Child component.</h1>

    <p>{{.counter}}</p>
    <button live-click="click">+</button>

    {{if (gt .counter 5)}}
        <h1>counter is greater then 5</h1>
    {{end}}
</div>