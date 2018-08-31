{{with .recovery_codes -}}
    <h1>Recovery codes regenerated</h1>
    <p>
        <span>Recovery Codes:</span></br>
        {{range . -}}
        <span>{{.}}</span><br />
        {{end -}}
    </p>
{{else -}}
    <span>{{.n_recovery_codes}} recovery codes remaining.</span>
    <form action="{{mountpathed "2fa/recovery/regen"}}" method="POST">
        {{with .error}}{{.}}<br />{{end}}
        {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
        <button type="submit">Regenerate</button>
    </form>
{{end -}}
