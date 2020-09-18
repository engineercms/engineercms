<!-- 视频列表 视频主页-->
<!DOCTYPE html>
<html lang="en">

<head>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/jquery.form.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css" />
  <script type="text/javascript" src="/static/js/select2.js"></script>
</head>

<body>
  <!-- <ul class="nav nav-tabs" id="myTab" role="tablist">
    <li class="nav-item" role="presentation">
      <a class="nav-link active" id="home-tab" data-toggle="tab" href="#home" role="tab" aria-controls="home" aria-selected="true">待审批</a>
    </li>
    <li class="nav-item" role="presentation">
      <a class="nav-link" id="profile-tab" data-toggle="tab" href="#profile" role="tab" aria-controls="profile" aria-selected="false">历史</a>
    </li>
    <li class="nav-item" role="presentation">
      <a class="nav-link" id="contact-tab" data-toggle="tab" href="#contact" role="tab" aria-controls="contact" aria-selected="false">Contact</a>
    </li>
  </ul> -->
  <div class="col-lg-12">
    <ul id="myTab" class="nav nav-tabs">
      <li class="active">
        <a href="#home" data-toggle="tab">
          待审批
        </a>
      </li>
      <li><a href="#ios" data-toggle="tab">历史</a></li>
    </ul>
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane fade in active" id="home">
        <div id="toolbar" class="btn-group">
          <button {{if .IsAdmin}} style="display:none" {{end}} type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-envelope-o"> 申请</i>
          </button>
          <button {{if not .IsAdmin}} style="display:none" {{end}} type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-share"> 分享码</i>
          </button>
          <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
            <i class="fa fa-trash"> 删除</i>
          </button>
        </div>
        <table id="table"></table>
      </div>
      <div class="tab-pane fade" id="ios">
        <table id="table1"></table>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $('#table').bootstrapTable({
    url: '/v1/cart/getusercart',
    search: 'true',

    showSearchClearButton: 'true',

    showRefresh: 'true',
    showColumns: 'true',
    toolbar: '#toolbar',
    pagination: 'true',
    sidePagination: "server",
    queryParamsType: '',
    //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
    // limit, offset, search, sort, order 否则, 需要包含: 
    // pageSize, pageNumber, searchText, sortName, sortOrder. 
    // 返回false将会终止请求。
    pageSize: 15,
    pageNumber: 1,
    pageList: [15, 50, 100],
    uniqueId: "id",
    // singleSelect:"true",
    clickToSelect: "true",
    showExport: "true",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val(),
        status: "0"
      };
      //搜索框功能
      //当查询条件中包含中文时，get请求默认会使用ISO-8859-1编码请求参数，在服务端需要对其解码
      // if (null != searchText) {
      //   try {
      //     searchText = new String(searchText.getBytes("ISO-8859-1"), "UTF-8");
      //   } catch (Exception e) {
      //     e.printStackTrace();
      //   }
      // }
      return param;
    },
    columns: [{
      title: '选择',
      // radio: 'true',
      checkbox: 'true',
      width: '10',
      align: "center",
      valign: "middle"
    }, {
      // field: 'Number',
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      align: "center",
      valign: "middle"
    }, {
      field: 'usernickname',
      title: '用户',
      align: "center",
      valign: "middle"
    }, {
      field: 'producttitle',
      title: '成果名称',
      halign: "center",
      valign: "middle"
    }, {
      field: 'projecttitle',
      title: '目录名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'topprojecttitle',
      title: '项目名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'status',
      title: '状态',
      formatter:'StatusFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'updated',
      title: '日期',
      formatter: 'localDateFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'action',
      title: '操作',
      formatter: 'actionFormatter',
      events: 'actionEvents',
      align: "center",
      valign: "middle"
    }]
  })

  $('#table1').bootstrapTable({
    url: '/v1/cart/getusercart',
    search: 'true',

    showSearchClearButton: 'true',

    showRefresh: 'true',
    showColumns: 'true',
    // toolbar: '#toolbar',
    pagination: 'true',
    sidePagination: "server",
    queryParamsType: '',
    //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
    // limit, offset, search, sort, order 否则, 需要包含: 
    // pageSize, pageNumber, searchText, sortName, sortOrder. 
    // 返回false将会终止请求。
    pageSize: 15,
    pageNumber: 1,
    pageList: [15, 50, 100],
    uniqueId: "id",
    // singleSelect:"true",
    clickToSelect: "true",
    showExport: "true",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val(),
        status: "1"
      };
      //搜索框功能
      //当查询条件中包含中文时，get请求默认会使用ISO-8859-1编码请求参数，在服务端需要对其解码
      // if (null != searchText) {
      //   try {
      //     searchText = new String(searchText.getBytes("ISO-8859-1"), "UTF-8");
      //   } catch (Exception e) {
      //     e.printStackTrace();
      //   }
      // }
      return param;
    },
    columns: [{
      title: '选择',
      // radio: 'true',
      checkbox: 'true',
      width: '10',
      align: "center",
      valign: "middle"
    }, {
      // field: 'Number',
      title: '序号',
      formatter: function(value, row, index) {
        return index + 1
      },
      align: "center",
      valign: "middle"
    }, {
      field: 'usernickname',
      title: '用户',
      align: "center",
      valign: "middle"
    }, {
      field: 'producttitle',
      title: '成果名称',
      halign: "center",
      valign: "middle"
    }, {
      field: 'projecttitle',
      title: '目录名称',
      align: "center",
      valign: "middle"
    }, {
      field: 'status',
      title: '状态',
      formatter:'StatusFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'updated',
      title: '日期',
      formatter: 'localDateFormatter',
      align: "center",
      valign: "middle"
    }, {
      field: 'topprojecttitle',
      title: '项目名称',
      align: "center",
      valign: "middle"
    }]
  })

  // 申请

  // 批量删除
  $("#deleteButton").click(function(e, value, row, index) {
    var selectRow = $('#table').bootstrapTable('getSelections');
    if (selectRow.length <= 0) {
      alert("请先勾选成果！");
      return false;
    }
    if (confirm("确定删除吗？")) {
      var title = $.map(selectRow, function(row) {
        return row.producttitle;
      })
      var ids = "";
      for (var i = 0; i < selectRow.length; i++) {
        if (i == 0) {
          ids = selectRow[i].id;
        } else {
          ids = ids + "," + selectRow[i].id;
        }
      }
      // var removeline=$(this).parents("tr")
      //提交到后台进行修改数据库状态修改
      $.ajax({
        type: "post", //这里是否一定要用post？？？
        url: "/v1/cart/deleteusercart",
        // data: JSON.stringify(selectRow3), //JSON.stringify(row),
        data: { ids: ids },
        success: function(data, status) { //数据提交成功时返回数据
          $('#table').bootstrapTable('remove', {
            field: 'producttitle',
            values: title
          });
          // removeline.remove();
          alert("删除“" + data.data + "”成功！(status:" + status + ".)");
          // $('#table1').bootstrapTable('refresh', {url:'/admin/merit/meritlist/1'});
        }
      });
    }
  })

  function actionFormatter(value, row, index) {
    return [
      '<a class="btn btn-danger btn-sm" href="javascript:void(0)" title="删除">',
      '<i id="delete" class="fa fa-trash-o"> Delete</i>',
      '</a>'
    ].join('');
  }

  window.actionEvents = {
    'click #delete': function(e, value, row, index) {
      if (confirm("确定删除吗？")) {
        var ids = "";
        ids = row.id
        // 删除前端表格用
        var mycars = new Array()
        mycars[0] = row;
        ids2 = $.map(mycars, function(row) {
          return row.id;
        });
        // var removeline=$(this).parents("tr")
        //提交到后台进行修改数据库状态修改
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: "/v1/cart/deleteusercart",
          data: { ids: ids },
          success: function(data, status) { //数据提交成功时返回数据
            $('#table').bootstrapTable('remove', {
              field: 'id',
              values: ids2
            });
            // removeline.remove();
            alert("删除“" + data.data + "”成功！(status:" + status + ".)");
            // $('#table1').bootstrapTable('refresh', {url:'/admin/merit/meritlist/1'});
          }
        });
      }
    }
    // if(confirm("确定删除吗？")){  
    // var removeline=$(this).parents("tr")
    //     $.ajax({
    //     type:"post",//这里是否一定要用post？？？
    //     url:"/admin/merit/deletemeritlist",
    //     data: {CatalogId:row.Id},
    //         success:function(data,status){//数据提交成功时返回数据
    //         removeline.remove();
    //         alert("删除“"+data+"”成功！(status:"+status+".)");
    //         }
    //     });  
    // }
  };

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD HH:mm:ss').add(8, 'hours').format('YYYY-MM-DD HH:mm:ss');
  }

  function StatusFormatter(value, row, index) {
    // alert(row.Status);
    if (row.status == 0) { //Status
      return '待审批';
    } else if (row.status == 1) {
      return '已审批';
    } else {
      return '失效';
    }
  }
  </script>
</body>