<template>
  <div class="app-container">
    <div style="height: 635px; text-align: center;padding: 20px 0; overflow: auto">
      <el-button style="width:200px;margin-top: 20px" @click="handleOpenQuestionnaire">
        <el-icon>
          <search/>
        </el-icon>
        Select File
      </el-button>
      <Questionnaire
          :question-list="questionList"
          :editable="false"
      ></Questionnaire>
    </div>
  </div>
</template>

<script>
import Questionnaire from '~/components/Questionnaire/index.vue'

export default {
  name: "ViewQuestionnaire",
  components: {Questionnaire},
  data() {
    return {
      questionList: []
    }
  },
  methods: {
    async handleOpenQuestionnaire() {
      try {
        this.questionList = await window.go.main.App.OpenQuestionnaire() ?? []
      } catch (e) {
        this.$message.error(e.message)
      }
    }
  }
}
</script>

<style scoped>

</style>