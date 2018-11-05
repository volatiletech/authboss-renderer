<h1>Authorize adding 2fa to your account</h1>
<form action="{{.url}}" method="POST">
    <input type="text" class="form-control" name="code" placeholder="E-mail" disabled="true" autocomplete="off" value="{{.email}}"><br />
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    <button type="submit">Ok</button>
</form>
