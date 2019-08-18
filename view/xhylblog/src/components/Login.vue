<template>
<div class="login">
    <HeadNav></HeadNav>
    <Card class="card">
        <Row>
            <span style="margin:10px;margin-left:15px;font-size: 15px;"><b>登录</b></span>
        </Row>
        <Row>
            <hr  style="height:1px;border:none;border-top:1px solid #e8eaec;margin:10px;margin-top:15px;"/>
        </Row>
        <Row>
            <div style="margin-top:5px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">账号</span></div>
        </Row>
        <Row>
            <Input placeholder="账号" clearable style="margin-left:15px;width:90%;" size="large" maxlength="20">
                <Icon type="ios-contact" slot="prefix" />
            </Input>
        </Row>
        <Row>
            <div style="margin-top:18px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">密码</span></div>
        </Row>
        <Row>
            <Input placeholder="密码"  style="margin-left:15px;width:90%;" :type="isLaws?'text':'password'" size="large" maxlength="16">
                <Icon type="ios-lock" slot="prefix" />
                <Icon :type="isLaws?'ios-eye-off':'ios-eye'" slot="suffix" @click="changeLaws" />
            </Input>
        </Row>
        <Row>
            <Button v-if="!openVerify" type="info" style="margin:15px;width:90%;" @click="changeOpenVerify" long><Icon type="md-arrow-round-down" />验证码</Button>
            <div v-else-if="openVerify" align="center" style="margin:10px;">
                <slide-verify :l="42"
                    :r="10"
                    :w="310"
                    :h="155"
                    @success="onSuccess"
                    @fail="onFail"
                    @refresh="onRefresh"
                    :slider-text="text"
                    ></slide-verify>
            </div>
            
        </Row>
        <Row>
            <Button type="success" style="margin:15px;width:90%;" :disabled="!isVerify" long>登录</Button>
        </Row>
        <Row>
            <div style="font-size: 14px;color:rgb(53, 184, 245);text-align:center;margin-top:15px;"><router-link to='/forgetpwd' tag="span" exact><span>忘记密码？</span></router-link></div>
            <div style="font-size: 14px;text-align:center;margin-top:5px;">尚未拥有账户？ <router-link to='/forgetpwd' tag="span" exact><span  style="color:rgb(53, 184, 245);">注册</span></router-link></div>
        </Row>
        <Row>
            <div style="margin-top:25px;margin:5px;margin-left:20px;">
                <span style="color:rgb(53, 184, 245);">快捷登录>></span>
                <img src="@/assets/qq.png" style="width:20px;height:20px;margin:2px;margin-bottom:10px;"/>
                <img src="@/assets/weixin.png" style="width:20px;height:20px;margin:2px;margin-bottom:10px;"/>
                <img src="@/assets/github.png" style="width:20px;height:20px;margin:2px;margin-bottom:10px;"/>
            </div>
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
                openVerify:false,
                isVerify: false,
                isLaws: false,
                text: '向右滑',
            }
        },
		methods:{
            onSuccess(){
            this.isVerify = true;
            },
            onFail(){
                this.isVerify = false;
            },
            onRefresh(){
                this.isVerify = false;
            },
            changeOpenVerify(){
                this.openVerify=true;
            },
            changeLaws(){
                if(this.isLaws){
                    this.isLaws=false;
                }else{
                    this.isLaws=true;
                }
                
            }
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
  width: 23%;
  margin:0 auto;
  margin-top:10%;
}



</style>