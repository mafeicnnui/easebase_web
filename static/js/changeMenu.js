let menu_change = {
    data() {
        return {
            menu_name: '',
            tableData: [],
            currentPage: 1, // 当前页码
            total: 20,      // 总条数
            pageSize: 10,   // 每页的数据条数,
            dialogFormVisible : false,
            dialogTypeFlag:'',
            dialogFormTitle : '',
            editForm: {
                id:'',
                parent:'',
                dm_parent:[],
                name: '',
                url: '',
                icon:'',
                status: '',
            }
        }
    },
    template: `
        <el-form  label-width="120px">
              <el-form-item label=""> </el-form-item>
              <el-row type="flex">
                 <el-col :span="24">
                      <el-form-item label="菜单名">                                        
                          <el-input placeholder="请输入菜单名" v-model="menu_name" @input="changeValue" style="width:400px"></el-input>
                          <el-button type="primary" @click="queryMenu">查询</el-button>
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
                      label="权限ID"
                      width="80">
              </el-table-column>
              <el-table-column                      
                      prop="Name"
                      label="菜单名">
                      <template slot-scope="scope">
                          <div v-html="scope.row.Name"></div>
                       </template>
              </el-table-column>
              <el-table-column                      
                      prop="Icon"
                      label="图标名">
              </el-table-column>
              <el-table-column
                      prop="Url"
                      label="功能链接">
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
                      prop="Status"
                      label="菜单状态"
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
                              @click="delMenu(scope.$index, scope.row)">删除</el-button>

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
                          <el-form-item label="菜单名称">                                        
                              <el-input placeholder="请输入菜单名" v-model="editForm.name" style="width:300px"></el-input>
                          </el-form-item>                      
                     </el-col>      
                     <el-col :span="12">
                          <el-form-item label="功能链接">                     
                               <el-input placeholder="请输入功能链接" v-model="editForm.url" style="width:300px"></el-input>
                          </el-form-item>                      
                      </el-col>
                 </el-row>
                 <el-row>                  
                     <el-col :span="12">
                        <el-form-item label="菜单状态">
                          <el-select v-model="editForm.status"  style="width:300px">
                                  <el-option label="启用" value="1"></el-option>
                                  <el-option label="禁用" value="0"></el-option>
                                </el-select>
                        </el-form-item>  
                      </el-col>          
                      <el-col :span="12">
                          <el-form-item label="菜单图标">
                             <el-input placeholder="请输入菜单图标" v-model="editForm.icon"  style="width:300px"></el-input>
                          </el-form-item>
                      </el-col>                  
                 </el-row>            
                 <el-row>
                   <el-col :span="12">
                     <el-form-item label="上级节点">
                        <el-select v-model="editForm.parent" placeholder="请选择上级节点" style="width:300px">
                            <el-option v-for="dm in editForm.dm_parent"  :key="dm.id" :label="dm.name" :value="dm.id"></el-option>
                        </el-select>
                     </el-form-item>   
                    </el-col> 
                  </el-row>
                  <el-form-item label=""> </el-form-item>   
               </el-form>    
                <div slot="footer" class="dialog-footer">
                      <el-button @click="dialogFormVisible = false">取 消</el-button>
                      <el-button type="primary" @click="updMenu">更 新</el-button>
                </div>
             </el-dialog>
        </el-form>`,
    methods:{
        getParent() {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/menu/parent'.format(server_ip,server_port),
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {

                    this.editForm.dm_parent = res.data['Data']
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
            this.queryMenu()
        },
        queryMenu() {
            axios({
                method: 'get',
                url: 'http://{0}:{1}/menu'.format(server_ip,server_port),
                params: {
                    name  : this.menu_name,
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
            this.dialogFormTitle = '菜单变更'
            this.itemReadOnly = true
            this.itemShow = true
            this.dialogTypeFlag = 'edit'
            //获取菜单信息
            axios({
                method: 'get',
                url: 'http://{0}:{1}/menu/{2}'.format(server_ip,server_port,row.Id),
                timeout: 1000,
            }).then((res) => {
                console.log('openEdit->res:',res)
                if (res.data['Code'] == 200 ) {
                    this.editForm.id        = res.data['Data'].Id
                    this.editForm.name      = res.data['Data'].Name
                    this.editForm.url       = res.data['Data'].Url
                    this.editForm.icon      = res.data['Data'].Icon
                    this.editForm.status    = res.data['Data'].Status
                    this.editForm.parent    = res.data['Data'].ParentId
                }
            }).catch((error) => {
                console.log('error=',error);
            });
        },
        openDetail:function(index,row){
            console.log('openDetail=',index,row)
        },
        updMenu:function(){
            axios({
                method: 'Post',
                url: 'http://{0}:{1}/menu'.format(server_ip,server_port),
                params: {
                    id          : this.editForm.id,
                    name        : this.editForm.name,
                    url         : this.editForm.url,
                    icon        : this.editForm.icon,
                    status      : this.editForm.status,
                    parent      : this.editForm.parent
                },
                timeout: 1000,
            }).then((res) => {
                if (res.data['Code'] == 200 ) {
                    this.$notify({
                        title: '成功',
                        message: '更新成功',
                        type: 'success'
                    });
                    this.queryMenu();
                }
            }).catch((error) => {
                console.log('error=',error);
            });

        },
        delMenu:function(index,row) {
            console.log("delMenu=",index,row)
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                axios.delete('http://{0}:{1}/menu/'.format(server_ip,server_port),
                    { params:
                            {
                                id:row.Id
                            }
                    }).then((res) => {
                    if (res.data['Code'] == 200 ) {
                        this.$notify({
                            title: '成功',
                            message: '菜单['+row.Name+']删除成功',
                            position: 'top-right',
                            type: 'success'
                        });
                        this.queryMenu();
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
        this.getParent();
        this.queryMenu();
    }

}