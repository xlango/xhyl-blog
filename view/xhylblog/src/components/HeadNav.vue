<template>
  <div id="head-nav">
    <div class="container-fluid dev-header" >
      <div class="row">
        <div class="col-lg-1"></div>
        <div class="col-lg-1"><img style="width:130px;height:50px;" src="@/assets/xhyl.png"></div>
        <div class="col-lg-4"></div>
        <div class="col-lg-4">
          <div class="container" style="font-size: 16px;color:rgba(40, 40, 40, 0.7);">
            <div class="row">
              <div class="col-md-1"><router-link to='/' tag="div" exact><div>首页</div></router-link></div>
              <div class="col-md-1"><router-link to='/article/add' tag="div" exact><div>文章</div></router-link></div>
              <div class="col-md-1"><router-link to='/question/list' tag="div" exact><div>提问</div></router-link></div>
              <div class="col-md-1"><router-link to='/b' tag="div" exact><div>相册</div></router-link></div>
              <div class="col-md-1"><router-link to='/b' tag="div" exact><div>归档</div></router-link></div>
            </div>
          </div>
        </div>
        <!-- <div class="col-md-1">文章</div>
        <div class="col-md-1">提问</div>
        <div class="col-md-1">相册</div> -->
        <div class="col-lg-2">
          <div class="container" style="font-size: 14px;color:rgb(53, 184, 245);">
            <div class="row">
              <div v-if="!isLogin" class="col-md-2">
                <router-link to='/login' tag="span" exact><span>登录</span></router-link>
                <b style="color:#000;">•</b>
                <router-link to='/registe' tag="span" exact><span>注册</span></router-link>
              </div>
              <div v-else-if="isLogin" class="col-md-2" align="left">
                <Dropdown>
                  <a href="javascript:void(0)">
                      <img src="@/assets/avatar.png" style="width:35px;height:35px;margin:5px;border-radius:70%;"/>
                  </a>
                  <DropdownMenu slot="list">
                      <DropdownItem><router-link to='/me' tag="span" exact>主页</router-link></DropdownItem>
                      <DropdownItem><div  @click="loginOut">注销</div></DropdownItem>
                  </DropdownMenu>
                </Dropdown>
                <Poptip title="提醒" content="！" placement="bottom-end">
                  <img src="@/assets/alert.png" style="width:30px;height:30px;margin:5px;"/>
                </Poptip>
                

                <Dropdown>
                  <a href="javascript:void(0)">
                      <img src="@/assets/add.png" style="width:25px;height:25px;margin-left:5px;"/>
                  </a>
                  <DropdownMenu slot="list">
                      <DropdownItem><router-link to='/circle/add' tag="span" exact>写文章</router-link></DropdownItem>
                      <DropdownItem><router-link to='/question/add' tag="span" exact>提问题</router-link></DropdownItem>
                      <DropdownItem><router-link to='/photo/add' tag="span" exact>传相册</router-link></DropdownItem>
                      <DropdownItem><router-link to='/archive/add' tag="span" exact>归档</router-link></DropdownItem>
                  </DropdownMenu>
                </Dropdown>
                
              </div>
            </div>
          </div>
        </div>
      </div>
</div>
  </div>
</template>
<script>
import HeadNav from "@/components/HeadNav";
export default {
  name: "head-nav",
  data() {
    return {
      isLogin: false
    };
  },
  methods: {
    loginOut() {
      this.$Spin.show({
        render: h => {
          return h("div", [
            h("Icon", {
              class: "demo-spin-icon-load",
              props: {
                type: "ios-loading",
                size: 18
              }
            }),
            h("div", "正在注销...")
          ]);
        }
      });
      setTimeout(() => {
        this.$Spin.hide();
        localStorage.removeItem("token");
      this.$router.push({ path: "/login" });
      }, 2000);
    }
  },
  created() {
    var token = localStorage.getItem("token");
    if (token) {
      this.isLogin = true;
    }
  }
};
</script>

<style scoped>
.dev-header {
  width: 100%;
  height: 50px;
  padding: 0 40px;
  box-sizing: border-box;
  z-index: 1000;
  position: fixed;
  top: 0;
  background-color: #fff;
  box-shadow: 0 1px 3px rgba(26, 26, 26, 0.1);
  font-size: 14px;
  line-height: 50px; /*设置其文字内容垂直居中*/
}
.router-link-active {
  color: rgb(0, 0, 0);
  text-decoration: underline;
}
.demo-spin-icon-load {
  animation: ani-demo-spin 1s linear infinite;
}
</style>