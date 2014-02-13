<!--Begin Datatables-->
<div class="row">
  <div class="col-lg-12">
    <div class="box">
      <header>
        <div class="icons">
          <i class="fa fa-table"></i>
        </div>
        <h5>用户列表</h5>
      </header>
      <div id="collapse4" class="body">
        <table id="dataTable" class="table table-bordered table-condensed table-hover table-striped">
          <thead>
            <tr>
              <th>用户ID</th>
              <th>用户名</th>
              <th>昵称</th>
              <th>Email</th>
              <th>备注</th>
              <th>状态</th>
              <th>最后登录时间</th>
            </tr>
          </thead>
          <tbody>
          {{range $key,$val:=.users}}
            <tr>
              <td>{{$val.Id}}</td>
              <td>{{$val.Username}}</td>
              <td>{{$val.Nickname}}</td>
              <td>{{$val.Email}}</td>
              <td>{{$val.Remark}}</td>
              <td>{{$val.Status}}</td>
              <td>{{$val.Lastlogintime}}</td>
            </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div><!-- /.row -->

<!--End Datatables-->
<script src="//ajax.googleapis.com/ajax/libs/jqueryui/1.10.3/jquery-ui.min.js"></script>
<script src="/static/metis/lib/datatables/jquery.dataTables.js"></script>
<script src="/static/metis/lib/datatables/DT_bootstrap.js"></script>
<script src="/static/metis/lib/tablesorter/js/jquery.tablesorter.min.js"></script>
<script src="/static/metis/lib/touch-punch/jquery.ui.touch-punch.min.js"></script>
 <script>
  $(function() {
    metisTable();
    metisSortable();
  });
</script>