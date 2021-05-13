let user_query = {
    data() {
        return {
            user_name: '',
            tableData: [],
            currentPage: 1, // 当前页码
            total: 20,      // 总条数
            pageSize: 10,   // 每页的数据条数,
        }
    },
    template: `
        <el-form  label-width="120px">
              <el-form-item label=""> </el-form-item>
              <el-row type="flex">
                 <el-col :span="24">
                      <el-form-item label="姓名">                                        
                          <el-input placeholder="请输入登陆名" v-model="user_name" @input="changeValue" style="width:400px"></el-input>
                          <el-button type="primary" @click="queryUser">查询</el-button>
                      </el-form-item>  
                 </el-col>
             </el-row>
             <el-table
                  :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
                  :header-cell-style="{'text-align':'center'}"
                  :cell-style="{'text-align':'center'}"
                  border
                  style="width: 100%"
                  :default-sort="{prop: 'EmpNo', order: 'descending'}" >
              <el-table-column
                      prop="EmpNo"
                      label="工号"
                      width="80">
              </el-table-column>
              <el-table-column
                      prop="Name"
                      label="姓名"
                      width="120">
              </el-table-column>
              <el-table-column
                      prop="LoginName"
                      label="用户名"
                      width="150">
              </el-table-column>
              <el-table-column
                      prop="Gender"
                      label="性别"
                      width="60">
              </el-table-column>
              <el-table-column
                      prop="Phone"
                      label="手机号"
                      width="120">
              </el-table-column>
              <el-table-column
                      prop="Email"
                      sortable
                      width="200"
                      label="电子邮箱">
              </el-table-column>
              <el-table-column
                      prop="DeptNo"
                      sortable
                      label="部门名" 
                      width="120">
              </el-table-column> 
              <el-table-column
                      prop="ExpireDate"
                      label="过期时间"
                      width="120">
              </el-table-column>
              <el-table-column
                      prop="Creator"
                      label="创建人"
                      width="100">
              </el-table-column>
              <el-table-column
                      prop="CreateDate"
                      label="创建时间"
                      width="180">
              </el-table-column>
              <el-table-column
                      prop="Updater"
                      label="更新人"
                      width="100">
              </el-table-column>         
              <el-table-column
                      prop="LastUpdateDate"
                      label="最近更新时间"
                      width="180">
              </el-table-column>
               <el-table-column
                      prop="Status"
                      label="状态"
                      width="80">
              </el-table-column>
          </el-table>
          <div class="block" style="margin-top:15px;">
              <el-pagination
                      align='center'
                      @size-change="handleSizeChange"
                      @current-change="handleCurrentChange"
                      :current-page="currentPage"
                      :page-sizes="[1,5,10,20]"
                      :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
                      :total="tableData.length">
              </el-pagination>
          </div>
          
        </el-form>`,
    methods:{
        handleSizeChange(val) {
            console.log(`每页 ${val} 条`);
            this.currentPage = 1;
            this.pageSize = val;
        },
        handleCurrentChange(val) {
            console.log(`当前页: ${val}`);
            this.currentPage = val;
        },
        changeValue: function() {
            this.queryUser()
        },
        queryUser() {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/user'.format(server_ip,server_port),
                params: {
                    name  : this.user_name,
                },
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    let newArr = res.data['Data'];
                    for (let i=0;i<newArr.length;i++)  {
                        newArr[i]['ExpireDate'] = newArr[i]['ExpireDate'].substr(0,10)
                    }
                    this.tableData = newArr
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
    },
    mounted: function() {
        this.queryUser();
    }

}