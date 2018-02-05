{{with .flash_success}}{{.}}{{end}}
<form action="{{mountpathed "login"}}" method="POST">
    {{with .error}}{{.}}<br />{{end}}
    <input type="text" class="form-control" name="email" placeholder="E-mail" value="{{.primaryIDValue}}"><br />
    <input  type="password" class="form-control" name="password" placeholder="Password"><br />
    <input type="hidden" name="{{.csrfName}}" value="{{.csrfToken}}" />
    {{if .showRemember}}<input type="checkbox" name="rm" value="true"> Remember Me{{end}}
    <button type="submit">Login</button><br />
    {{if .showRecover}}<a href="{{mountpathed "recover"}}">Recover Account</a>{{end}}
    {{if .showRegister}}<a href="{{mountpathed "register"}}">Register Account</a>{{end}}
</form>
