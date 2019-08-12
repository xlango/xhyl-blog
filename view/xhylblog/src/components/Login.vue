<template>
<div class="login">
    <HeadNav></HeadNav>
    <Card class="card">
        <Row>
            <span style="margin:10px;font-size: 15px;"><b>登录</b></span>
        </Row>
        <Row>
            <hr  style="height:1px;border:none;border-top:1px solid rgb(136, 135, 135);margin:10px;margin-top:15px;"/>
        </Row>
        <Row>
            <div style="margin:10px;margin-top:5px;"><span style="color:red;">*</span><span style="margin-left:5px;">账号</span></div>
        </Row>
    </Card>
</div>
</template>
<script>
import HeadNav from '@/components/HeadNav'
export default {
  name: "login",
  components:{
            HeadNav,
        },
  data () {
            return {
                
            }
        },
		methods:{

		},
		created(){
			
		}
};
</script>
<style scoped>
.login {
  height: 100%;
}
.card {
  width: 25%;
  height: 500px;
  margin:0 auto;
  margin-top:10%;
}



</style>