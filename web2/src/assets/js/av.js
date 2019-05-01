import AV from 'leancloud-storage'
import {Message} from 'element-ui'
import {i18n} from "@/i18n"


export const uploadAvatar = function (file, name, callback) {
    let cloudFile = new AV.File(name, file)
    cloudFile.save().then(function (res) {
        callback(res)
    }, function (err) {
        console.log(err)
        Message.error(i18n.t("message.profile.UAF"))
    })
}

export default AV