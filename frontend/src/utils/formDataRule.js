export function requiredRule(tip = '不能为空') {
    return {
        required: true,
        trigger: ["blur", "input"],
        message: tip
    }
}

export function regexpRule(regexp = (str) => {
    return true
}, tip = '格式错误') {
    return {
        required: true,
        trigger: ["blur", "change"],
        validator: (rule, value) => {
            if (!value) {
                return new Error('请输入')
            }
            if (!regexp(value)) {
                return new Error(tip)
            }
            return true
        }
    }
}