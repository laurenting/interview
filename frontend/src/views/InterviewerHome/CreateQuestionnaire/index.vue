<template>
  <div class="app-container">
    <div style="height: 635px; padding: 20px 0; overflow: auto">
      <el-form
          ref="createQuestionnaireForm"
          :model="dynamicValidateForm"
          label-width="120px"
      >
        <el-form-item label="" label-width="40px">
          <h2>Make Question List</h2>
        </el-form-item>
        <el-form-item
            v-for="(q, index) in dynamicValidateForm.question"
            :key="q.key"
            :label="'Q' + (index + 1)"
            :prop="'question.' + index + '.value'"
            :rules="{
        required: true,
        message: 'question is empty',
        trigger: 'blur',
      }"
        >
          <el-input
              type="textarea"
              v-model="q.value"
              :autosize="{ minRows: 1, maxRows: 4 }"
              class="question-input"/>
          <el-button type="danger" @click.prevent="removeQuestion(q)">
            Delete
          </el-button
          >
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">Submit</el-button>
          <el-button type="success" @click="addQuestion">New Question</el-button>
          <el-button @click="resetForm">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CreateQuestionnaire',
  data() {
    return {
      dynamicValidateForm: {
        question: []
      }
    }
  },
  methods: {
    removeQuestion(q) {
      const index = this.dynamicValidateForm.question.indexOf(q)
      if (index !== -1) {
        this.dynamicValidateForm.question.splice(index, 1)
      }
    },
    addQuestion() {
      this.dynamicValidateForm.question.push({
        key: Date.now(),
        value: ''
      })
    },
    resetForm() {
      this.$refs.createQuestionnaireForm.resetFields()
    },
    async submitForm() {
      const questionnaire = this.dynamicValidateForm.question
      if (!questionnaire.length) {
        return this.$message.warning("Can't submit an empty questionnaire")
      }
      const result = await this.$refs['createQuestionnaireForm'].validate()
      if (!result) return
      try {
        const resp = await window.go.main.App.Save(this.dynamicValidateForm.question)
        if (!resp) return
        this.$message.success(resp)
      } catch (e) {
        this.$message.error(e.message)
      }
    }
  }
}
</script>

<style scoped>
.question-input {
  margin: 5px 10px 0 0;
  width: 500px;
  overflow: auto;
}
</style>
