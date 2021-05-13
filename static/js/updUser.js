let user_change = {
    data() {
        return {
            input1: '',
            input2: '',
            input3: '',
            input4: '',
            select:'',
            num1: 1,
            num2: 1,
            num3: 1,
            num4: 1,
        }
    },
    template: `
          <div>
              <el-input placeholder="请输入内容2" v-model="input1">
                <template slot="prepend">Http://</template>
              </el-input>
              <p></p>
               <el-input placeholder="请输入内容2" v-model="input2">
                 <template slot="prepend">Http://</template>
              </el-input>
              <p></p>
              <el-input placeholder="请输入内容" v-model="input3">
                <template slot="append">.com</template>
              </el-input>
              <p></p>
               <el-input-number v-model="num1"></el-input-number>
               <el-input-number size="medium" v-model="num2"></el-input-number>
               <el-input-number size="small" v-model="num3"></el-input-number>
               <el-input-number size="mini" v-model="num4"></el-input-number>
               <p></p>
               <el-input placeholder="请输入内容" v-model="input4" class="input-with-select">
                <el-select v-model="select" slot="prepend" placeholder="请选择">
                  <el-option label="餐厅名"   value="1"></el-option>
                  <el-option label="订单号"   value="2"></el-option>
                  <el-option label="用户电话" value="3"></el-option>
                </el-select>
                <el-button slot="append" icon="el-icon-search"></el-button>
              </el-input>
          </div>
      `,
    methods:{
        handleSelect(item) {
            console.log(item);
        },
    },
}