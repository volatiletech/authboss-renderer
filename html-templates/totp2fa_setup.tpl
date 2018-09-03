<h1>Setup two-factor authentication</h1>
<form action="{{mountpathed "2fa/totp/setup"}}" method="POST">
    <button type="submit">Begin Setup</button>
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
</form>
