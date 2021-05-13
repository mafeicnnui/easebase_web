let role_query = {
    data() {
        return {
            role_name: '',
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
                      <el-form-item label="角色名">                                        
                          <el-input placeholder="请输入角色名" v-model="role_name" @input="changeValue" style="width:400px"></el-input>
                          <el-button type="primary" @click="queryRole">查询</el-button>
                      </el-form-item>  
                 </el-col>
             </el-row>
             <el-table
                  :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
                  :header-cell-style="{'text-align':'center'}"
                  :cell-style="{'text-align':'center'}"
                  border
                  style="width: 100%;white-space:pre-wrap;"
                  :default-sort="{prop: 'EmpNo', order: 'descending'}" >
              <el-table-column
                      prop="Id"
                      label="角色ID"
                      width="80">
              </el-table-column>
              <el-table-column                      
                      prop="Name"
                      label="角色名" 
                      width="250">
              </el-table-column>
              <el-table-column
                      prop="Status"
                      label="是否启用"
                      width="120">
              </el-table-column>
              <el-table-column
                      prop="Creator"
                      label="创建人"
                      width="150">
              </el-table-column>
              <el-table-column
                      prop="CreateDate"
                      label="创建时间" >
              </el-table-column>
              <el-table-column
                      prop="Updater"
                      label="更新人"
                      width="150">
              </el-table-column>         
              <el-table-column
                      prop="LastUpdateDate"
                      label="最近更新时间" >
              </el-table-column>
              <el-table-column
                      fixed="right"
                      label="操作"
                      width="150">
                     <template slot-scope="scope">
                       <el-button
                          size="mini"
                          @click="openDetail(scope.$index, scope.row)">详情</el-button>
                     </template>
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
            this.queryRole()
        },
        changeValue: function() {
            this.queryRole()
        },
        openDetail:function(index,row){
            console.log('openDetail=',index,row)
        },
        queryRole() {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/role'.format(server_ip,server_port),
                params: {
                    name  : this.role_name,
                },
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    this.tableData =  res.data['Data'];
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
    },
    mounted: function() {
        this.queryRole();
    }

}