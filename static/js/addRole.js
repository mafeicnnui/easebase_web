let role_add = {
    data() {
        return {
            form:{
                name:'',
                status:'1',
                SysPrivileges: [],
                rolePrivileges: [],
                titles :['系统权限','角色权限'],
            }
        }
    },
    template: `
           <el-form :inline="true"  ref="form" :model="form" label-width="120px">
              <el-form-item label=""> </el-form-item>
              <el-row type="flex">
                 <el-col :span="24">
                      <el-form-item label="角色名称">                                        
                          <el-input placeholder="请输入角色名称" v-model="form.name" style="width:780px"></el-input>
                      </el-form-item>                      
                 </el-col>    
              </el-row>
              <el-row>      
                 <el-col :span="24">
                    <el-form-item label="角色状态">
                      <el-select v-model="form.status"  style="width:780px">
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
                                v-model="form.rolePrivileges"
                                :props="{
                                    key  : 'value',
                                    label: 'desc'                 
                                }"
                                :data="form.SysPrivileges"
                                :titles="form.titles"
                                :filterable="true"
                                >
                          </el-transfer>
                      </el-form-item>
                  </el-col>                  
              </el-row>   
               <el-form-item label=""> </el-form-item>              
               <el-row type="flex">              
                   <el-col :span="14"  style="text-align: center">
                      <el-form-item>
                        <el-button type="primary" @click="saveRole">保存</el-button>
                      </el-form-item>
                   </el-col>              
               </el-row>              
        </el-form>
      `,
    methods:{
        getQx() {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/menu/qx'.format(server_ip,server_port),
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                  this.form.SysPrivileges = res.data['Data']
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
        saveRole() {
            console.log('form:',this.form)
            axios({
                method: 'Put',
                url: 'http://{0}:{1}/role'.format(server_ip,server_port),
                params: {
                    name    : this.form.name,
                    status  : this.form.status,
                    role_privileges  : this.form.rolePrivileges,
                },
                timeout: 60000,
            }).then((res) => {
                console.log('saveRole:',res)
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
        },
        clearForm() {
            this.form.name =''
            this.form.status=''
            this.form.rolePrivileges=[]
            this.getQx();
        }
    },
    mounted: function() {
      this.getQx();
    } ,
}