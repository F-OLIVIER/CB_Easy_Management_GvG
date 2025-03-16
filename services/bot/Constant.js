// fichier annexe
import { logToFile } from "./log.js";

// Module nodejs et npm
import { Client, GatewayIntentBits, MessageFlags, Partials } from "discord.js";

// Clients Discord
export const client = new Client({
  intents: [
    GatewayIntentBits.Guilds, // Permet au bot de recevoir des événements liés aux serveurs (guilds).
    GatewayIntentBits.GuildMessages, // Permet au bot de recevoir les messages envoyés dans les salons textuels d'un serveur.
    GatewayIntentBits.MessageContent, // permet au bot de lire le contenu des messages envoyés dans les salons textuels.
    GatewayIntentBits.GuildMembers, // Permet au bot de recevoir les ajouts, suppressions et mises à jour des membres d'un serveur (⚠️ doit être activé sur le portail Discord).
    GatewayIntentBits.GuildVoiceStates, // Permet au bot de suivre les connexions, déconnexions et changements d'état des membres dans les salons vocaux.
  ],
  partials: [
    Partials.Message, // Permet d’accéder aux messages supprimés ou non mis en cache.
  ],
});

// Réponse d'interaction à l'utilisateur
export async function reponseUserInteraction(interaction, msg) {
  await interaction.reply({
    content: "<@" + interaction.user.id + ">\n" + msg,
    flags: MessageFlags.Ephemeral,
  });
}

// Log dans un chan Discord (chan utilisateur ou chan de gestion)
export function msgChanDiscord(ID_Group, ID_Chan, msg) {
  const chan = client.channels.cache.get(ID_Chan);
  if (!chan) {
    logToFile(`Chan ${ID_Chan} innexistant`, "errors_bot.log");
    return;
  }
  if (!chan.permissionsFor(client.user)?.has(["SendMessages", "AttachFiles", "EmbedLinks"])) {
    logToFile(`Le bot n'a pas la permission d'envoyer des messages dans ${ID_Chan}`, "errors_bot.log");
    return;
  }
  chan.send("<@&" + ID_Group + ">\n" + msg);
}

// Message utilisateur retirer de la database
export function UserLeave(ID_Chan_Gestion, name, nickname, msg) {
  const chan = client.channels.cache.get(ID_Chan_Gestion);
  if (!chan) {
    logToFile(`Chan ${ID_Chan_Gestion} innexistant`, "errors_bot.log");
    return;
  }
  if (!chan.permissionsFor(client.user)?.has(["SendMessages", "AttachFiles", "EmbedLinks"])) {
    logToFile(`Le bot n'a pas la permission d'envoyer des messages dans ${ID_Chan_Gestion}`, "errors_bot.log");
    return;
  }
  chan.send(name + " (" + nickname + ") " + msg);
}
