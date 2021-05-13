let user_add = {
    data() {
        return {
            form: {
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
        <el-form :inline="true"  ref="form" :model="form" label-width="120px">
              <el-form-item label=""> </el-form-item>
              <el-row type="flex">
                 <el-col :span="12">
                      <el-form-item label="姓名">                                        
                          <el-input placeholder="请输入姓名" v-model="form.user_name" style="width:550px"></el-input>
                      </el-form-item>                      
                 </el-col>      
                 <el-col :span="12">
                      <el-form-item label="用户">                     
                           <el-input placeholder="请输入用户名" v-model="form.login_name" style="width:550px"></el-input>
                      </el-form-item>                      
                  </el-col>
             </el-row>
             <el-row>                  
                 <el-col :span="12">
                    <el-form-item label="工号">
                      <el-input placeholder="请输入工号" v-model="form.emp_no" style="width:550px"></el-input>
                    </el-form-item>  
                  </el-col>          
                  <el-col :span="12">
                      <el-form-item label="密码">
                         <el-input placeholder="请输入密码" v-model="form.password" show-password style="width:550px"></el-input>
                      </el-form-item>
                  </el-col>                  
             </el-row>            
             <el-row>
               <el-col :span="12">
                 <el-form-item label="性别">
                    <el-select v-model="form.gender" placeholder="请选择性别" style="width:550px">
                        <el-option v-for="dm in form.dm_gender"  :key="dm.dmm" :label="dm.dmmc" :value="dm.dmm"></el-option>
                    </el-select>
                 </el-form-item>   
                </el-col>                    
                <el-col :span="12">
                 <el-form-item label="部门">
                    <el-select v-model="form.dept_no" placeholder="请选择部门" style="width:550px">
                       <el-option v-for="dm in form.dm_dept_no"  :key="dm.dmm" :label="dm.dmmc" :value="dm.dmm"></el-option>
                    </el-select>
                 </el-form-item>   
                </el-col>
              </el-row>
              <el-row>              
                  <el-col :span="12">
                     <el-form-item label="邮箱">
                      <el-input placeholder="请输入邮箱" v-model="form.email" style="width:550px"></el-input>
                     </el-form-item> 
                   </el-col>
                   <el-col :span="12">
                     <el-form-item label="手机">
                        <el-input placeholder="请输入手机" v-model="form.phone" style="width:550px"></el-input>
                     </el-form-item>
                  </el-col>             
              </el-row>
              <el-row>               
                  <el-col :span="12">
                     <el-form-item label="过期日期">
                        <el-date-picker 
                            type="date" 
                            placeholder="选择日期" 
                            v-model="form.expire_date" 
                            value-format="yyyy-MM-dd"
                            style="width: 550px">
                        </el-date-picker>
                    </el-form-item>
                  </el-col>                            
                  <el-col :span="12">
                       <el-form-item label="状态">
                            <el-select v-model="form.status"  style="width:550px">
                              <el-option label="启用" value="1"></el-option>
                              <el-option label="禁用" value="2"></el-option>
                            </el-select>
                       </el-form-item>     
                  </el-col>                  
               </el-row>
               <el-row>
                 <el-col :span="12">
                     <el-form-item label="角色">
                         <el-select v-model="form.user_role" multiple placeholder="请选择角色"  style="width:550px">
                            <el-option
                              v-for="item in form.sys_role"
                                  :key="item.value"
                                  :label="item.label"
                                  :value="item.value">
                            </el-option>
                        </el-select>
                    </el-form-item>
                  </el-col>               
               </el-row>    
               <el-form-item label=""> </el-form-item>              
               <el-row type="flex">              
                   <el-col :span="24"  style="text-align: center">
                      <el-form-item>
                        <el-button type="primary" @click="saveUser">保存</el-button>
                      </el-form-item>
                   </el-col>              
               </el-row>              
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
                        this.form.dm_gender = res.data['Data']
                    } else if (p_dm == '01') {
                        this.form.dm_dept_no = res.data['Data']
                    } else {

                    }
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
        saveUser() {
            axios({
                method: 'Put',
                url: 'http://{0}:{1}/user'.format(server_ip,server_port),
                params: {
                    name        : this.form.user_name,
                    emp_no      : this.form.emp_no,
                    gender      : this.form.gender,
                    email       : this.form.email,
                    dept_no     : this.form.dept_no,
                    expire_date : this.form.expire_date,
                    password    : this.form.password,
                    status      : this.form.status,
                    login_name  : this.form.login_name,
                    phone       : this.form.phone,
                    user_role   : this.form.user_role,
                },
                timeout: 1000,
            }).then((res) => {
                console.log('saveUser:',res)
                if (res.data['Code'] == 200 ) {
                    this.$notify({
                        title: '成功',
                        message: '保存成功',
                        type: 'success'
                    });
                    this.clearForm();
                }
            }).catch((error) => {
                console.log('error=',error);
            });

            console.log('form=',this.form);
            this.getDm();
        },
        clearForm() {
            this.form.login_name =''
            this.form.user_name='',
            this.form.emp_no='',
            this.form.password='',
            this.form.gender= '',
            this.form.dept_no='',
            this.form.email='',
            this.form.phone='',
            this.form.expire_date='',
            this.form.status='1',
            this.form.sys_role= [
                {value:1,label:'系统管理员'},
                {value:2,label:'开发人员'},
                {value:3,label:'测试人员'},
                {value:4,label:'实施人员'},
                {value:5,label:'数据库管理员'},
             ],
            this.form.user_role=[]
        }
    },
    mounted: function() {
        this.getDm('04')
        this.getDm('01')
    } ,
}