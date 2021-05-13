let user_change = {
    data() {
        return {
            user_name: '',
            tableData: [],
            currentPage: 1, // 当前页码
            total: 20,      // 总条数
            pageSize: 10,   // 每页的数据条数,
            dialogFormVisible : false,
            dialogTypeFlag:'',
            dialogFormTitle : '',
            editForm: {
                id:'',
                login_name: '',
                user_name: '',
                emp_no:'',
                password: '',
                gender: '',
                dm_gender: [],
                dept_no:'',
                dm_dept_no:[],
                email:'',
                phone:'',
                expire_date:'',
                status:'1',
                sys_role: [
                    {value:1,label:'系统管理员'},
                    {value:2,label:'开发人员'},
                    {value:3,label:'测试人员'},
                    {value:4,label:'实施人员'},
                    {value:5,label:'数据库管理员'},
                ],
                user_role: []
            }
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
                  style="width: 100%;"
                  border
                  :header-cell-style="{'text-align':'center'}"
                  :cell-style="{'text-align':'center'}"
                  :default-sort="{prop: 'EmpNo', order: 'descending'}" >
              <el-table-column
                      prop="EmpNo"
                      label="工号"
                      sortable
                      width="80">
              </el-table-column>
              <el-table-column
                      prop="Name"
                      label="姓名"
                      sortable
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
                      width="80">
              </el-table-column>
              <el-table-column
                      prop="Phone"
                      label="手机号"
                      width="120">
              </el-table-column>
              <el-table-column
                      prop="Email"
                      sortable
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
                      sortable
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
                      sortable
                      label="创建时间"
                      width="180">
              </el-table-column>
               <el-table-column
                      prop="Status"
                      label="状态"
                      width="80">
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
                              @click="delUser(scope.$index, scope.row)">删除</el-button>

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
                  <el-form :inline="true"  ref="editForm" :model="editForm" label-width="120px"> 
                      <el-form-item label=""> </el-form-item>
                      <el-row type="flex">
                         <el-col :span="12">
                              <el-form-item label="姓名">                                        
                                  <el-input placeholder="请输入姓名" v-model="editForm.user_name" style="width:300px"></el-input>
                              </el-form-item>                      
                         </el-col>      
                         <el-col :span="12">
                              <el-form-item label="用户">                     
                                   <el-input placeholder="请输入用户名" v-model="editForm.login_name" style="width:300px"></el-input>
                              </el-form-item>                      
                          </el-col>
                      </el-row>
                      <el-row>                  
                         <el-col :span="12">
                            <el-form-item label="工号">
                              <el-input placeholder="请输入工号" v-model="editForm.emp_no" style="width:300px"></el-input>
                            </el-form-item>  
                          </el-col>          
                          <el-col :span="12">
                              <el-form-item label="密码">
                                 <el-input placeholder="请输入密码" v-model="editForm.password" show-password style="width:300px"></el-input>
                              </el-form-item>
                          </el-col>                  
                      </el-row>            
                      <el-row>
                       <el-col :span="12">
                         <el-form-item label="性别">
                            <el-select v-model="editForm.gender" placeholder="请选择性别" style="width:300px">
                                <el-option v-for="dm in editForm.dm_gender"  :key="dm.dmm" :label="dm.dmmc" :value="dm.dmm"></el-option>
                            </el-select>
                         </el-form-item>   
                        </el-col>                    
                        <el-col :span="12">
                         <el-form-item label="部门">
                            <el-select v-model="editForm.dept_no" placeholder="请选择部门" style="width:300px">
                               <el-option v-for="dm in editForm.dm_dept_no"  :key="dm.dmm" :label="dm.dmmc" :value="dm.dmm"></el-option>
                            </el-select>
                         </el-form-item>   
                        </el-col>
                      </el-row>
                      <el-row>              
                          <el-col :span="12">
                             <el-form-item label="邮箱">
                              <el-input placeholder="请输入邮箱" v-model="editForm.email" style="width:300px"></el-input>
                             </el-form-item> 
                           </el-col>
                           <el-col :span="12">
                             <el-form-item label="手机">
                                <el-input placeholder="请输入手机" v-model="editForm.phone" style="width:300px"></el-input>
                             </el-form-item>
                          </el-col>             
                      </el-row>
                      <el-row>               
                          <el-col :span="12">
                             <el-form-item label="过期日期">
                                <el-date-picker 
                                    type="date" 
                                    placeholder="选择日期" 
                                    v-model="editForm.expire_date" 
                                    value-format="yyyy-MM-dd"
                                    style="width: 300px">
                                </el-date-picker>
                            </el-form-item>
                          </el-col>                            
                          <el-col :span="12">
                               <el-form-item label="状态">
                                    <el-select v-model="editForm.status"  style="width:300px">
                                      <el-option label="启用" value="1"></el-option>
                                      <el-option label="禁用" value="0"></el-option>
                                    </el-select>
                               </el-form-item>     
                          </el-col>                  
                      </el-row>
                      <el-row>
                         <el-col :span="24">
                             <el-form-item label="角色">
                                 <el-select v-model="editForm.user_role" multiple placeholder="请选择角色"  style="width:760px">
                                    <el-option
                                      v-for="item in editForm.sys_role"
                                          :key="item.value"
                                          :label="item.label"
                                          :value="item.value">
                                    </el-option>
                                </el-select>
                            </el-form-item>
                          </el-col>               
                       </el-row>   
                   </el-form>
            
                   <div slot="footer" class="dialog-footer">
                      <el-button @click="dialogFormVisible = false">取 消</el-button>
                      <el-button type="primary" @click="updUser">更新</el-button>
                  </div>
             </el-dialog>
          
        </el-form>`,
    methods:{
        getDm(p_dm) {
            axios({
                method: 'post',
                url: 'http://{0}:{1}/dm'.format(server_ip,server_port),
                params: {
                    dm: p_dm,
                },
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    if (p_dm == '04') {
                        this.editForm.dm_gender = res.data['Data']
                    } else if (p_dm == '01') {
                        this.editForm.dm_dept_no = res.data['Data']
                    } else {

                    }
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
        updUser:function(){
            axios({
                method: 'Post',
                url: 'http://{0}:{1}/user'.format(server_ip,server_port),
                params: {
                    id          : this.editForm.id,
                    name        : this.editForm.user_name,
                    emp_no      : this.editForm.emp_no,
                    gender      : this.editForm.gender,
                    email       : this.editForm.email,
                    dept_no     : this.editForm.dept_no,
                    expire_date : this.editForm.expire_date,
                    password    : this.editForm.password,
                    status      : this.editForm.status,
                    login_name  : this.editForm.login_name,
                    phone       : this.editForm.phone,
                    user_role   : this.editForm.user_role,
                },
                timeout: 10000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    this.$notify({
                        title: '成功',
                        message: '更新成功',
                        type: 'success'
                    });
                }
            }).catch((error) => {
                console.log('error=',error);
            });

        },
        openEdit:function(index,row){
            this.dialogFormVisible = true
            this.dialogFormTitle = '用户变更'
            this.itemReadOnly = true
            this.itemShow = true
            this.dialogTypeFlag = 'edit'
            //获取用户信息
            axios({
                method: 'get',
                url: 'http://{0}:{1}/user/{2}'.format(server_ip,server_port,row.id),
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    this.editForm.id          = res.data['Data'].Id
                    this.editForm.user_name   = res.data['Data'].Name
                    this.editForm.login_name  = res.data['Data'].LoginName
                    this.editForm.emp_no      = res.data['Data'].EmpNo
                    this.editForm.password    = res.data['Data'].Password
                    this.editForm.gender      = res.data['Data'].Gender
                    this.editForm.dept_no     = res.data['Data'].DeptNo
                    this.editForm.email       = res.data['Data'].Email
                    this.editForm.phone       = res.data['Data'].Phone
                    this.editForm.expire_date = res.data['Data'].ExpireDate.substr(0,10)
                    this.editForm.status      = res.data['Data'].Status

                    //获取用户角色信息
                    axios({
                        method: 'get',
                        url: 'http://{0}:{1}/user/role/{2}'.format(server_ip,server_port,res.data['Data'].Id),
                        timeout: 1000,
                    }).then((res) => {
                        if (res.data['Code'] == 200 ) {
                            let roles = [];
                            for (let i=0;i<res.data['Data'].length;i++)  {
                                roles[i]= res.data['Data'][i]['RoleId']
                            }
                            this.editForm.user_role  = roles
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
        delUser:function(index,row) {
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                axios.delete('http://{0}:{1}/user/'.format(server_ip,server_port),
                    { params:
                            {
                                id:row.id
                            }
                    }).then((res) => {
                        if (res.data['Code'] == 200 ) {
                            this.$notify({
                                title: '成功',
                                message: '用户['+row.Name+']删除成功',
                                position: 'top-right',
                                type: 'success'
                            });
                            this.queryUser();
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
        this.getDm('04');
        this.getDm('01');
        this.queryUser();
    }
}