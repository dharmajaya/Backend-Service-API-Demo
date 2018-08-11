<div class="container" style="padding-top: 80px" id="content">
  <div class="row">
    <form class="form-horizontal" method="POST">
       <div class="form-group{{if .Errors.AssetName}} has-error has-feedback{{end}}">
              <label for="Asset" class="col-sm-2 control-label">AssetName</label>
              <div class="col-sm-8">
                <input type="text" name="AssetName" id="AssetName" class="form-control" value="{{.Market.AssetName}}">
                {{if .Errors.AssetName}}<span class="help-block">{{.Errors.AssetName}}</span>{{end}}
              </div>
            </div>

            <div class="form-group{{if .Errors.NewPrice}} has-error has-feedback{{end}}">
              <label for="NewPrice" class="col-sm-2 control-label">NewPrice</label>
              <div class="col-sm-8">
                <input type="text" name="NewPrice" class="form-control" id="NewPrice" value="{{.Market.NewPrice}}">
                {{if .Errors.NewPrice}}<span class="help-block">{{.Errors.Last}}</span>{{end}}
              </div>
            </div>

            <div class="form-group{{if .Errors.BoardNews}} has-error has-feedback{{end}}">
              <label for="BoardNews" class="col-sm-2 control-label">BoardNews</label>
              <div class="col-sm-8">
                <input type="text" name="BoardNews" class="form-control" id="BoardNews" value="{{.Market.BoardNews}}">
                {{if .Errors.BoardNews}}<span class="help-block">{{.Errors.BoardNews}}</span>{{end}}
              </div>
            </div>

            <div class="form-group{{if .Errors.News}} has-error has-feedback{{end}}">
              <label for="News" class="col-sm-2 control-label">News</label>
              <ul><li>
              <div class="col-sm-8">
                <input type="text" name="News" class="form-control" id="News" value="{{.Market.News}}">
                {{if .Errors.News}}<span class="help-block">{{.Errors.News}}</span>{{end}}
              </div>
              </li>
              </ul>
            </div>

            <div class="form-group{{if .Errors.PromoteNews}} has-error has-feedback{{end}}">
              <label for="PromoteNews" class="col-sm-2 control-label">PromoteNews</label>
              <ul><li>
              <div class="col-sm-8">
                <input type="text" name="PromoteNews" class="form-control" id="PromoteNews" value="{{.Market.PromoteNews}}">
                {{if .Errors.PromoteNews}}<span class="help-block">{{.Errors.PromoteNews}}</span>{{end}}
              </div>
              </li>
              </ul>
            </div>

              <div class="form-group">
              <div class="col-sm-offset-2 col-sm-8">
                <button type="submit" class="btn btn-primary" value="save">Save</button>
              </div>
            </div>
    </form>
  </div>
  <hr>
  <div class="panel panel-danger">
    <div class="panel-heading">
      <h3 class="panel-title">Danger zone</h3>
    </div>
  </div>
</div><!-- end container -->
