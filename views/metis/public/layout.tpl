{{template "../public/header.tpl"}}
    <script src="/static/metis/lib/jquery.min.js"></script>
    <script src="/static/metis/lib/bootstrap/js/bootstrap.min.js"></script>
    <!--<script type="text/javascript" src="/static/metis/js/style-switcher.js"></script>-->
    <script src="/static/metis/js/main.min.js"></script>
  <body>
    <div id="wrap">
      <!-- #top -->
      {{template "../public/top.tpl"}}
      <!-- /#top -->
      <!-- #left-->
      {{template "../public/left.tpl"}}
      <!-- /#left -->
      <div id="content">
        <div class="outer">
          <div class="inner">
            {{.LayoutContent}}
          </div>
          <!-- end .inner -->
        </div>
        <!-- end .outer -->
      </div>
      <!-- end #content -->
    </div><!-- /#wrap -->
    <div id="footer">
      <p>2013 &copy; Metis Admin</p>
    </div>
    <!-- #helpModal -->
    {{template "../public/help.tpl"}}
    <!-- /#helpModal -->
  </body>
</html>