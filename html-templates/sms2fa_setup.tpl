<h1>Setup two-factor authentication</h1>
<form action="{{mountpathed "2fa/sms/setup"}}" method="POST">
    <input type="text" class="form-control" name="phone_number" placeholder="Phone Number" {{with .sms_phone_number}}value="{{.}}"{{end}}/>
    <button type="submit">Begin Setup</button>
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
</form>
