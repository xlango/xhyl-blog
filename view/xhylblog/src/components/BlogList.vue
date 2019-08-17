<template>
<div class="blog-list">
   <div class="container-fluid">
        <div class="row">
            <div class="col-md-1"></div>
            <div class="col-md-7">
                <div class="container-fluid">
                    <div class="row">
                        <Card class="card container">
                            <div class="row">
                                <div class="col-md-8">
                                    <img src="@/assets/logo.png" style="width:20px;height:20px;margin:5px;border-radius:70%;"/>
                                    <span style="margin:5px">张三</span>
                                    <span style="margin:5px">2019-01-01 18:00:00</span>
                                </div>
                                <div class="col-md-4"><Icon type="ios-megaphone-outline" style="float:right;"/></div>
                            </div>
                            <div class="row">
                                <h4 class="col-md-8" style="margin:7px;font-weight:bold;">博客标题</h4>
                            </div>
                            <div class="row">
                                <div class="col-md-8"><span style="margin:9px;">预览内容。</span></div>
                                <div class="col-md-4"><img src="@/assets/logo.png" style="width:80px;height:80px;margin:5px;float:right;"/></div>
                            </div>
                            <div class="row">
                                <div class="col-md-11">
                                    <button class="btntype"><img src="@/assets/logo.png" style="width:13px;height:13px;margin:2px;"/>vue</button>
                                    <button class="btntype"><img src="@/assets/logo.png" style="width:13px;height:13px;margin:2px;"/>vue</button>
                                </div>
                                <div class="col-md-1" style="color:green;margin-top:15px;" align="right" title="共回复2条">
                                    <Icon type="ios-chatbubbles" /><span style="font-size: 12px;margin:2px;" >2</span>
                                </div>
                                
                            </div>
                        </Card>
                        <Card class="card">
                            <div style="text-align:center">
                                <div>

                                </div>
                                
                            </div>
                        </Card>
                        <Card class="card">
                            <div style="text-align:center">
                                
                                
                            </div>
                        </Card>
                        <Button class="btnmore" title="点击加载更多"><Icon type="md-arrow-down" /></Button>
                    </div>
                </div>
                
            </div>
            <div class="col-md-3">
                <Card class="card-slider">
                    <Input suffix="ios-search" placeholder="搜索" style="width: 100%" />
                </Card>
                <Card class="card-slider">
                    <DatePicker :value="nowTime" format="yyyy年MM月dd日" type="date" placeholder="Select date" style="width: 100%;"></DatePicker>
                </Card>
                <Card class="card-slider">
                    <Row>
                        <Icon type="ios-options" /><span style="margin:5px;">分类：</span>
                    </Row>
                    <Row>
                        <hr  style="height:1px;border:none;border-top:1px solid rgb(136, 135, 135);margin:10px;"/>
                    </Row>
                    <Row>
                        <Tag checkable color="primary">Vue</Tag>
                        <Tag checkable color="success">Java</Tag>
                        <Tag checkable color="error">C#</Tag>
                        <Tag checkable color="warning">Golang</Tag>
                    </Row>
                </Card>
                
                <Button type="success"  class="card-slider" long><Icon type="ios-create-outline"/>写文章</Button>
            </div>
            <div class="col-md-1"></div>
        </div>
    </div>
</div>
</template>
<script>
import BlogListCard from "@/components/BlogListCard";

export default {
  name: "blog-list",
  data () {
            return {
                nowTime:new Date().toLocaleString(),
            }
        },
		methods:{

		},
		mounted(){

		}
};
</script>
<style scoped>
.blog-list {
  height: 100%;
}
.card {
  width: 90%;
  height: 220px;
  margin-top: 20px;
  margin-left: 90px;
}
.card-slider {
  width: 80%;
  margin-top: 20px;
}

.btntype {
  height: 23px;
  width: 65px;
  margin: 5px;
  display: inline-block;
  font-size: 14px;
  font-weight: 400;
  text-align: center;
  background-image: none;
  color: #333;
  background-color: white;
    color: black;
    border: 1px solid #e7e7e7;
}
.btnmore {
  width: 90%;
  margin-top: 20px;
  margin-left: 90px;
}

</style>