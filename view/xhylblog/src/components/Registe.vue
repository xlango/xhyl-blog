<template>
<div class="registe">
    <HeadNav></HeadNav>
    <Card class="card">
        <Tabs value="name1">
            <TabPane label="普通注册" name="name1">
                <Row>
                    <div style="margin-top:5px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">账号</span></div>
                </Row>
                <Row>
                    <Input v-model="username" placeholder="账号" clearable style="margin-left:15px;width:90%;" size="large" maxlength="20">
                        <Icon type="ios-contact" slot="prefix" />
                    </Input>
                </Row>
                <Row>
                    <div style="margin-top:18px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">密码</span></div>
                </Row>
                <Row>
                    <Input v-model="password" placeholder="密码"  style="margin-left:15px;width:90%;" :type="isLaws?'text':'password'" size="large" maxlength="16">
                        <Icon type="ios-lock" slot="prefix" />
                        <Icon :type="isLaws?'ios-eye-off':'ios-eye'" slot="suffix" @click="changeLaws" />
                    </Input>
                </Row>
                <Row>
                    <div style="margin-top:18px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">确认密码</span></div>
                </Row>
                <Row>
                    <Input v-model="surepassword" placeholder="确认密码"  style="margin-left:15px;width:90%;" :type="isLaws?'text':'password'" size="large" maxlength="16">
                        <Icon type="ios-lock" slot="prefix" />
                        <Icon :type="isLaws?'ios-eye-off':'ios-eye'" slot="suffix" @click="changeLaws" />
                    </Input>
                </Row>
                <Row>
                    <div style="margin-top:5px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">昵称</span></div>
                </Row>
                <Row>
                    <Input v-model="name" type="text" placeholder="昵称" clearable style="margin-left:15px;width:90%;" size="large" maxlength="20">
                        <Icon type="ios-bookmark" slot="prefix" />
                    </Input>
                </Row>
                <Row>
                    <div style="margin-top:5px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">邮箱</span></div>
                </Row>
                <Row>
                    <Input v-model="email" type="text" placeholder="邮箱" clearable style="margin-left:15px;width:90%;" size="large" maxlength="20">
                        <Icon type="ios-mail" slot="prefix" />
                    </Input>
                </Row>
                <Row>
                    <div style="margin-top:5px;margin:5px;margin-left:15px;"><span style="color:red;">*</span><span style="margin-left:5px;">手机</span></div>
                </Row>
                <Row>
                    <Input v-model="tel" type="text" placeholder="手机" clearable style="margin-left:15px;width:90%;" size="large" maxlength="20">
                        <Icon type="ios-call" slot="prefix" />
                    </Input>
                </Row>
                <Row>
                    <Input v-model="verifyCode" type="text" placeholder="验证码" clearable style="margin-left:15px;margin-top:10px;width:50%;" size="large" maxlength="20"/>
                    <Button :size="buttonSize" type="info" style="margin-left:5px;margin-top:10px;">获取验证码</Button>
                </Row>
                <Row>
                    <Checkbox v-model="single" style="margin-left:15px;margin-top:10px;">同意遵守浪哥相关协议</Checkbox>
                    <Button type="success" style="margin-left:15px;margin-top:10;width:90%;"  @click="registe" long>注册</Button>
                </Row>
                <Row>
                    <div style="font-size: 14px;text-align:center;margin-top:5px;">已拥有账户？ <router-link to='/login' tag="span" exact><span  style="color:rgb(53, 184, 245);">立即登录</span></router-link></div>
                </Row>
            </TabPane>
            <TabPane label="快捷注册" name="name2">
                
            </TabPane>
        </Tabs>
        

    </Card>
</div>
</template>
<script>
import HeadNav from '@/components/HeadNav'

export default {
  name: "registe",
  components:{
            HeadNav,
        },
        data () {
            return {
                username:"",
                password:"",
                name:"",
                email:"",
                tel:"",
                verifyCode:"",
                surepassword:"",
                isLaws: false,
            }
        },
		methods:{
            //登录
            registe(){
                this.$post("/user/register", {
                    username:this.username,
                    password:this.password,
                    name:this.name,
                    email:this.email,
                    tel:this.tel,
                }).then(
                    res => {
                        if (res.Code == 200) {
                            alert("注册成功");
                        }
                    },
                    err => {

                    }
                );
            },
            changeLaws(){
                if(this.isLaws){
                    this.isLaws=false;
                }else{
                    this.isLaws=true;
                }
            },
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
  margin:0 auto;
  margin-top:10%;
}



</style>