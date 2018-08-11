<div class="container" style="padding-top: 80px" id="content">
  <div class="row">
    <form class="form-horizontal" method="POST">
      <div class="form-group{{if .Errors.Reference_Link}} has-error has-feedback{{end}}">
              <ul><li>
              <label for="Reference_Link" class="col-sm-2 control-label">Reference_Link </label>
              <div class="col-sm-8">
                <input type="text" name="Reference_Link" id="Reference_Link" class="form-control" value="{{.User.Reference_Link}}">
                {{if .Errors.Reference_Link}}<span class="help-block">{{.Errors.Reference_Link}}</span>{{end}}
              </div>
              </li></ul>
            </div>

            <div class="form-group{{if .Errors.Registers_count}} has-error has-feedback{{end}}">
            <ul><li>
              <label for="Registers_count" class="col-sm-2 control-label">Registers_count</label>
              <div class="col-sm-8">
                <input type="int" name="Registers_count" class="form-control" id="Registers_count" value="{{.User.Registers_count}}">
                {{if .Errors.Registers_count}}<span class="help-block">{{.Errors.Registers_count}}</span>{{end}}
              </div>
              </li></ul>
            </div>

            <div class="form-group{{if .Errors.Click_count}} has-error has-feedback{{end}}">
              <label for="Click_count" class="col-sm-2 control-label">Click_count</label>
              <div class="col-sm-8">
                <input type="int" name="Click_count" class="form-control" id="Click_count" value="{{.User.Click_count}}">
                {{if .Errors.Click_count}}<span class="help-block">{{.Errors.Click_count}}</span>{{end}}
              </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-8">
                <button type="submit" class="btn btn-primary" value="save">save</button>
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
