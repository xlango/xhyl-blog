<template>
<div class="home">
    <HeadNav></HeadNav>
    <BlogList class="bloglist"></BlogList>
</div>
</template>
<script>
import HeadNav from '@/components/HeadNav'
import BlogList from '@/components/BlogList'
    export default {
        name:"home",
        components:{
            HeadNav,
            BlogList
        }
    }
</script>
<style scoped>
.bloglist{
    margin-top:50px;
    height:100%;
}
</style>