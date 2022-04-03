<template>
  <div>
    <el-menu router class="el-menu-demo" mode="horizontal">
      <el-menu-item index="/">
        <el-icon><arrow-left-bold/></el-icon>
        Back
      </el-menu-item>
      <el-menu-item index="/interviewer/create_questionnaire">
        <el-icon><document-add /></el-icon>
        Create Questionnaire
      </el-menu-item>
      <el-menu-item index="/interviewer/view_questionnaire" >
        <el-icon><document-checked /></el-icon>
        Import Questionnaire
      </el-menu-item>
    </el-menu>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'InterviewerHome',
  components: {},

  data () {
    return {}
  },
  methods: {}
}
</script>

<style scoped>

</style>
