<div id="left">
  <!-- #menu -->
  <ul id="menu" class="collapse">
    <li class="active">
      <a href="javascript:;">菜单<span class="fa arrow"></span> </a>
      <ul>
      {{range $key,$val:=.tree}}
        <li class="active">
          <a href="javascript:;">{{$val.Text}}<span class="fa arrow"></span> </a>
          <ul class="in" style="height: auto;">
          {{range $k,$v:=$val.Children}}
            <li> <a href="{{$v.Attributes.Url}}">{{$v.Text}}</a> </li>
          {{end}}
          </ul>
        </li>
        {{end}}
      </ul>
    </li>
  </ul><!-- /#menu -->
</div>