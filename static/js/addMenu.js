let menu_add = {
    data() {
        return {
            form: {
                name: '',
                url: '',
                status:'',
                parent: '',
                dm_parent: [],
                icon: '',
            }
        }
    },
    template: `
        <el-form :inline="true"  ref="form" :model="form" label-width="120px">
              <el-form-item label=""> </el-form-item>
              <el-row type="flex">
                 <el-col :span="12">
                      <el-form-item label="菜单名称">                                        
                          <el-input placeholder="请输入菜单名" v-model="form.name" style="width:550px"></el-input>
                      </el-form-item>                      
                 </el-col>      
                 <el-col :span="12">
                      <el-form-item label="功能链接">                     
                           <el-input placeholder="请输入功能链接" v-model="form.url" style="width:550px"></el-input>
                      </el-form-item>                      
                  </el-col>
             </el-row>
             <el-row>                  
                 <el-col :span="12">
                    <el-form-item label="菜单状态">
                      <el-select v-model="form.status"  style="width:550px">
                              <el-option label="启用" value="1"></el-option>
                              <el-option label="禁用" value="2"></el-option>
                            </el-select>
                    </el-form-item>  
                  </el-col>          
                  <el-col :span="12">
                      <el-form-item label="菜单图标">
                         <el-input placeholder="请输入菜单图标" v-model="form.icon" show-password style="width:550px"></el-input>
                      </el-form-item>
                  </el-col>                  
             </el-row>            
             <el-row>
               <el-col :span="12">
                 <el-form-item label="上级节点">
                    <el-select v-model="form.parent" placeholder="请选择上级节点" style="width:550px">
                        <el-option v-for="dm in form.dm_parent"  :key="dm.id" :label="dm.name" :value="dm.id"></el-option>
                    </el-select>
                 </el-form-item>   
                </el-col> 
              </el-row>
               <el-form-item label=""> </el-form-item>              
               <el-row type="flex">              
                   <el-col :span="24"  style="text-align: center">
                      <el-form-item>
                        <el-button type="primary" @click="saveMenu">保存</el-button>
                      </el-form-item>
                   </el-col>              
               </el-row>              
        </el-form>`,
    methods:{
        getParent() {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/menu/parent'.format(server_ip,server_port),
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {

                    this.form.dm_parent = res.data['Data']
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
        saveMenu() {
            axios({
                method: 'Put',
                url: 'http://{0}:{1}/menu'.format(server_ip,server_port),
                params: {
                    name: this.form.name,
                    url: this.form.url,
                    status: this.form.status,
                    icon: this.form.icon,
                    parent: this.form.parent,
                },
                timeout: 1000,
            }).then((res) => {
                console.log('saveMenu:',res)
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
            this.getParent();
        },
        clearForm() {
            this.form.name =''
            this.form.url='',
            this.form.status='',
            this.form.icon='',
            this.form.parent= ''
        }
    },
    mounted: function() {
        this.getParent()
    } ,
}