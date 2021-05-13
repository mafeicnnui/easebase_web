let role_change = {
    data() {
        return {
            role_name: '',
            tableData: [],
            currentPage: 1, // 当前页码
            total: 20,      // 总条数
            pageSize: 10,   // 每页的数据条数,
            dialogFormVisible : false,
            dialogTypeFlag:'',
            dialogFormTitle : '',
            editForm: {
                name:'',
                status:'1',
                SysPrivileges: [],
                rolePrivileges: [],
                titles :['系统权限','角色权限'],
            },
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
                          width="250">
                      <template slot-scope="scope">
                          <el-button
                                  size="mini"
                                  @click="openDetail(scope.$index, scope.row)">详情</el-button>
                          <el-button
                                  size="mini"
                                  @click="openEdit(scope.$index, scope.row)">编辑</el-button>
                          <el-button
                                  size="mini"
                                  type="danger"
                                  @click="delRole(scope.$index, scope.row)">删除</el-button>
    
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
          
              <el-dialog :title="dialogFormTitle" :visible.sync="dialogFormVisible">
                <el-form :inline="true"  ref="editForm" :model="editForm" label-width="80px">
                      <el-form-item label=""> </el-form-item>
                      <el-row type="flex">
                         <el-col :span="24">
                              <el-form-item label="角色名称">                                        
                                  <el-input placeholder="请输入角色名称" v-model="editForm.name" style="width:780px"></el-input>
                              </el-form-item>                      
                         </el-col>    
                      </el-row>
                      <el-row>      
                         <el-col :span="24">
                            <el-form-item label="角色状态">
                              <el-select v-model="editForm.status"  style="width:780px">
                                      <el-option label="启用" value="1"></el-option>
                                      <el-option label="禁用" value="0"></el-option>
                                    </el-select>
                            </el-form-item>  
                          </el-col>     
                     </el-row>
                     <el-row>         
                          <el-col :span="24">
                              <el-form-item label="角色权限">
                                 <el-transfer
                                        v-model="editForm.rolePrivileges"
                                        :props="{
                                            key  : 'value',
                                            label: 'desc'                 
                                        }"
                                        :data="editForm.SysPrivileges"
                                        :titles="editForm.titles"
                                        :filterable="true"
                                        >
                                  </el-transfer>
                              </el-form-item>
                          </el-col>                  
                      </el-row>   
                </el-form> 
                <div slot="footer" class="dialog-footer">
                      <el-button @click="dialogFormVisible = false">取 消</el-button>
                      <el-button type="primary" @click="updRole">更 新</el-button>
                </div>
             </el-dialog>
        </el-form>`,
    methods:{
        getQx(p_roleId) {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/menu/role/qx'.format(server_ip,server_port),
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    this.editForm.SysPrivileges = res.data['Data']
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
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
        openEdit:function(index,row){
            console.log('openEdit->row1:',row)
            this.dialogFormVisible = true
            this.dialogFormTitle = '角色变更'
            this.itemReadOnly = true
            this.itemShow = true
            this.dialogTypeFlag = 'edit'
            //获取角色信息
            axios({
                method: 'get',
                url: 'http://{0}:{1}/role/{2}'.format(server_ip,server_port,row.Id),
                timeout: 1000,
            }).then((res) => {
                console.log('openEdit->res:',res)
                if (res.data['Code'] == 200 ) {
                    this.editForm.id        = res.data['Data'].Id
                    this.editForm.name      = res.data['Data'].Name
                    this.editForm.status    = res.data['Data'].Status
                    this.getQx(this.editForm.id);

                    //获取角色权限信息
                    axios({
                        method: 'get',
                        url: 'http://{0}:{1}/role/privileges/{2}'.format(server_ip,server_port,res.data['Data'].Id),
                        timeout: 3000,
                    }).then((res) => {
                        console.log('openEdit->res2:',res)
                        if (res.data['Code'] == 200 ) {
                            let privileges =[]
                            for (let i=0;i<res.data['Data'].length;i++)  {
                                privileges[i] = res.data['Data'][i]['value']
                            }
                            this.editForm.rolePrivileges= privileges
                        }

                    }).catch((error) => {
                        console.log('error=',error);
                    });

                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
        openDetail:function(index,row){
            console.log('openDetail=',index,row)
        },
        updRole:function(){
            console.log("updRole:",this.editForm)
            axios({
                method: 'Post',
                url: 'http://{0}:{1}/role'.format(server_ip,server_port),
                params: {
                    id          : this.editForm.id,
                    name        : this.editForm.name,
                    status      : this.editForm.status,
                    role_privileges  : this.editForm.rolePrivileges,
                },
                timeout: 3000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    this.$notify({
                        title: '成功',
                        message: '更新成功',
                        type: 'success'
                    });
                    this.queryRole();
                }
            }).catch((error) => {
                console.log('error=',error);
            });

        },
        delRole:function(index,row) {
            console.log("delMenu=",index,row)
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                axios.delete('http://{0}:{1}/role/'.format(server_ip,server_port),
                    { params:
                            {
                                role_id:row.Id
                            }
                    }).then((res) => {
                    if (res.data['Code'] == 200 ) {
                        this.$notify({
                            title: '成功',
                            message: '角色['+row.Name+']删除成功',
                            position: 'top-right',
                            type: 'success'
                        });
                        this.queryRole();
                    } else {
                        this.$notify.error({
                            title: '错误',
                            message: res.data['Msg'],
                            position: 'top-right',
                            type: 'success'
                        });
                    }
                }).catch((error) => {
                    console.log('error=',error);
                });

            }).catch((error) => {
                console.log('error=',error);
                this.$message({
                    type: 'info',
                    message: '已取消删除',
                    position: 'top-right',
                    type: 'success'
                });
            });
        },
    },
    mounted: function() {
        this.queryRole();

    }

}