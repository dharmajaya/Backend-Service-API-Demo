<div class="container" style="padding-top: 80px" id="content">
  <div class="row">
    <form class="form-horizontal" method="POST">
      <div class="form-group{{if .Errors.UserName}} has-error has-feedback{{end}}">
        <label for="UserName" class="col-sm-2 control-label">UserName</label>
        <div class="col-sm-8">
          <input type="text" name="UserName" id="UserName" class="form-control" value="{{.User.UserName}}">
          {{if .Errors.UserName}}<span class="help-block">{{.Errors.UserName}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.NickName}} has-error has-feedback{{end}}">
        <label for="NickName" class="col-sm-2 control-label">NickName</label>
        <div class="col-sm-8">
          <input type="text" name="NickName" class="form-control" id="NickName" value="{{.User.NickName}}">
          {{if .Errors.NickName}}<span class="help-block">{{.Errors.NickName}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Email}} has-error has-feedback{{end}}">
        <label for="Email" class="col-sm-2 control-label">Email address</label>
        <div class="col-sm-8">
          <input type="email" name="Email" class="form-control" id="Email" value="{{.User.Email}}">
          {{if .Errors.Email}}<span class="help-block">{{.Errors.Email}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Phone1}} has-error has-feedback{{end}}">
        <label for="Phone" class="col-sm-2 control-label">Main phone</label>
        <div class="col-sm-8">
          <input type="tel" name="Phone" class="form-control" id="Phone" value="{{.User.Phone}}" placeholder="Example: 647-123-4567">
          {{if .Errors.Phone}}<span class="help-block">{{.Errors.Phone}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Phone2}} has-error has-feedback{{end}}">
        <label for="Phone2" class="col-sm-2 control-label">Secondary phone</label>
        <div class="col-sm-8">
          <input type="tel" name="Phone2" class="form-control" id="Phone2" value="{{.User.Phone2}}" placeholder="Example: 647-123-4567">
          {{if .Errors.Phone2}}<span class="help-block">{{.Errors.Phone2}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Location}} has-error has-feedback{{end}}">
        <label for="Location" class="col-sm-2 control-label">Location</label>
        <div class="col-sm-8">
          <input type="text" name="Location" class="form-control" id="Location" value="{{.User.Location}}">
          {{if .Errors.Location}}<span class="help-block">{{.Errors.Location}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.CurrentPassword}} has-error has-feedback{{end}}">
        <label for="CurrentPassword" class="col-sm-2 control-label">Current password</label>
        <div class="col-sm-8">
          <input type="password" name="CurrentPassword" class="form-control" id="CurrentPassword">
          {{if .Errors.CurrentPassword}}<span class="help-block">{{.Errors.CurrentPassword}}</span>{{end}}
        </div>
      </div>

      <div class="form-group">
        <div class="col-sm-offset-2 col-sm-8">
          <button type="submit" class="btn btn-primary" value="Update">Update</button>
        </div>
      </div>

    </form>
  </div>
  <hr>
  <div class="panel panel-danger">
    <div class="panel-heading">
      <h3 class="panel-title">Danger zone</h3>
    </div>
    <div class="panel-body">
      <p class="text-center"><a class="btn btn-danger" href="//{{.base_url}}/user/delete">Remove account</a></p>
    </div>
  </div>
</div><!-- end container -->
