<template>
<div class="archive">
    <HeadNav></HeadNav>
   <div class="container-fluid" style="margin-top: 50px;">
        <div class="row">
            <div class="col-md-1"></div>
            <div class="col-md-7">
                <div class="container-fluid">
                    <div class="row">
                        <Card class="card container">
                            <Timeline>
                                <TimelineItem>
                                    <p class="time">1976年</p>
                                    <p class="content">Apple I 问世</p>
                                </TimelineItem>
                                <TimelineItem>
                                    <p class="time">1984年</p>
                                    <p class="content">发布 Macintosh</p>
                                </TimelineItem>
                                <TimelineItem>
                                    <p class="time">2007年</p>
                                    <p class="content">发布 iPhone</p>
                                </TimelineItem>
                                <TimelineItem>
                                    <p class="time">2010年</p>
                                    <p class="content">发布 iPad</p>
                                </TimelineItem>
                                <TimelineItem>
                                    <p class="time">2011年10月5日</p>
                                    <p class="content">史蒂夫·乔布斯去世</p>
                                </TimelineItem>
                            </Timeline>
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
import HeadNav from "@/components/HeadNav";

export default {
    name: "archive",
    components:{
            HeadNav,
        },
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
.archive {
  height: 100%;
}
.card {
  width: 90%;
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
.time{
        font-size: 14px;
        font-weight: bold;
    }
    .content{
        padding-left: 5px;
    }
</style>